# simplecache

This is a pure in-memory cache for backends

This allows for classic CRUD operations

There is no in-disk data storage

This is a cache, not a message broker or a hybrid mix with in-disk data. The concept of a cache is to stay in-memory, since the DB already stores data in-disk.

`SCL` is `Simplecache Caching Language`, the language used to manage operations inside the cache.

## Run locally

To run:
 1. Clone via `git clone`
 2. Run `go mod tidy`
 3. Start by running `go run main.go <command>`

## Run locally with parser/lexer
 1. Follow steps 1,2 from above
 2. run `cd pkg/scl && goyacc -o scl_parser.go -p yy scl.y`
 3. run `cd ../.. && go run .`

Alternatively, you can compile and install the go package via `go install`, and then reference commands without `go run main.go`

Default port is 4000 but running with `-p <port>` allows for custom port.

## Concepts

The system is thought to be simple, with effective cache invalidation.

### Set data

To add data to the cache, use `SET <collection>.<key>:<value>` where value can be anything. From any default type (int, string, bool) to more complex maps, json, binaries

There are some options such as TTI (Time To Invalidate) as to when invalidate a record. If the cache gets a record invalidated, when being queried, the cache returns no hit (`nil`)


### Get data

To get data from the cache, run `GET <collection>.<key>` and the data will be returned. Since the data is anything, it will be returned as an `interface{} | any`.

This also means that you can get all the data from a collection running `GET <collection>.*`. This returns `[]interface{} | []any`.

If, as above, the record was invalidated, the operation will return an empty record (`nil`).


### Delete data
To delete data, run `DELETE <collection>.<key>`. You can delete all the data from a collection with `TRUNCACE | DROP <collection>`.

You can delete a single data element via key, or drop the dedicated collection:
 - Drop will unlink the collection directly, so it becomes unavailable to query;
 - Truncate will keep the collection but delete every record in it by traversing the tree and deleting each element avoiding tree-rebalancing (this locks the table).


### Update data
To update data, run `UPDATE <collection>.<key>:<new_value>`

The new value will override the old data directly. If the relation `<collection>.<key>` returns no record, an insert operation (SET) is performed.

Via update, you can also change the TTI of the record, by running `UPDATE <collection>.<key> TTI=<new_time_to_invalidate>`. If this relation has no record, a SET is performed with TTI