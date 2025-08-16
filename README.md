# TempoDB

A temporal-driven NoSQL versioned database.

While still working on it, this projects aims at answering the question:
- What happens to the data that is unused after a long time?

Let's give an example. Say we have a Postgres instance of all the matches to be ever played in NBA, from the first ever match on November the 1st 1946, until today. This totals up to over 60000 matches played.

If there are tables related to statistics for each player for each match, how likely is it that someone wants to find out the details from a match in 1950? Very unlikely.

For this reason, when a certain time threshold is met, data should be moved away from the disk the DB resides on and be put in a 'cold storage'.

What would be the benefits of this architecture?
 1. First off, removes data from storage. Data that is never read is backed up in a cold storage since it's never accessed and is just sat there.
 2. Second of which, this simplifies the LSMT structure of the DB, as less leaves are kept in the system and this increases the speed the tree is traversed, as well as efficiency of a bloom filter.


On top of this, since this is a NoSQL database (like CassandraDB), the system keep data in memory until moved to storage based on threshold. This will be improved with the addon of rules, which are specific configurations that allow the TempoDB to expose data in-memory (instad of disk) directly from a table, with specific columns only.

Now, versioning and replication.

This DB uses a similar SQL-Like language with some updates and changes. Moreover, it allows for table versioning with either Semantic (semver) or Incremental Versioning.

Replication for updates never touch data that is put in cold storage, since updating that data makes no sense under the idea that is never accessed.

However, when there is data in-memory and an update is made, the data is replicated based on the configuration of the platform:
- loose replication -> data is updated in-memory and it's considerd a success. Then, the data is replicated in the background on-disk
if the replication on-disk fails, the data is lost.
- strong replication -> data must be saved both in-memory and on-disk for the system to consider it a success. This is mroe consistency but slower