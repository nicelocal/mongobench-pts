package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RunBench(out io.Writer, coll *mongo.Collection, name string, fn func(int, int, *mongo.Collection, *sync.WaitGroup), cnt int, multithread bool) {
	var mt string
	var threads int
	if multithread {
		mt = "multi-threaded"
		threads = runtime.NumCPU()
	} else {
		mt = "single-threaded"
		threads = 1
	}

	fmt.Fprintf(out, "\nStarting %s: %d times %s\n", name, cnt, mt)
	var wg sync.WaitGroup
	t := time.Now()

	for t := range threads {
		wg.Add(1)
		go fn(cnt, t*cnt, coll, &wg)
	}
	wg.Wait()

	fmt.Fprintf(out, "Results for %s: %d times %s took %f seconds\n", name, cnt, mt, time.Since(t).Seconds())
}

func RunReadOne(cnt int, offset int, coll *mongo.Collection, wg *sync.WaitGroup) {
	d := bson.D{{Key: "idx", Value: ""}}
	for i := range cnt {
		d[0].Value = fmt.Sprintf("Key %d", offset+i)
		f := coll.FindOne(context.Background(), d)
		if f == nil {
			panic("Could not find!")
		}
	}
	wg.Done()
}

type TestWrite struct {
	Name string
	Idx  string
}

func RunWrite(cnt int, offset int, coll *mongo.Collection, wg *sync.WaitGroup) {
	v := TestWrite{}
	for i := range cnt {
		v.Idx = fmt.Sprintf("Value %d", i+offset)
		v.Idx = fmt.Sprintf("Key %d", i+offset)
		_, err := coll.InsertOne(context.Background(), v)
		if err != nil {
			panic(err)
		}
	}
	wg.Done()
}

func main() {
	opts := options.Client().ApplyURI("mongodb://127.0.0.1")

	mongoSession, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal(err)
	}

	coll := mongoSession.Database("test").Collection("test")
	coll.Drop(context.Background())
	err = coll.Database().CreateCollection(context.Background(), "test")
	if err != nil {
		panic(err)
	}

	var out io.Writer = os.Stdout
	logFile := os.Getenv("LOG_FILE")
	if logFile != "" {
		f, err := os.Create(logFile)
		if err != nil {
			log.Fatal("Could not open logfile:", err)
		}
		out = io.MultiWriter(out, f)
	}

	for multithread, cnt := range map[bool]int{false: 100000, true: 100000} {
		indexModel := mongo.IndexModel{
			Keys: bson.D{{"idx", 1}},
		}
		_, err = coll.Indexes().CreateOne(context.TODO(), indexModel)
		if err != nil {
			panic(err)
		}

		RunBench(out, coll, "write", RunWrite, cnt, multithread)

		RunBench(out, coll, "read", RunReadOne, cnt, multithread)

		err = coll.Drop(context.Background())
		if err != nil {
			panic(err)
		}
	}
}
