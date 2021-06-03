# Test assignment for Analytics Software Engineer position

This repo contains Github event data for 1 hour.

The goal is to write a script that outputs:

- Top 10 active users by amount of PRs created and commits pushed (username, PRs count, commits count)
- Top 10 repositories by amount of commits pushed (repo name, commits count)
- Top 10 repositories by amount of watch events (repo name, watch events count)

This assignment must be done in any type-safe language, that the candidate prefers.
---

### Here are my solutions
I have add two approach to solve the given problem

#### `git_events_in-memory.py`
Here I am trying the normal approach assuming I have enough memory to hold all the `key:value` in-memory.
In this approach I am parsing events and loading all `key:value` in memory and holding the data in-memory until I get all results
> This is not an efficient approach, depending the data-size and data-type and machine with not enough RAM, heap will grow and eventually we will OOM error.

---
#### `git_events_using_file.py`
This approach is slightly better then IN-Memory approach though we are reading all the data In-Memory but we are not keeping all the data
to RAM, after we get all the data we are writing all sorted `key:value` to disk on single disk write.

- Here we might also face the same OOM issue ( depending on Data-Type,Data-Size or running on a machine with not enough RAM ) give only for my
  case I parsed all the data to reduce the size.
- As we are writing all the sorted data on disk and the data is already sorted we just need to read top lines (this will reduce the no I/O call on disk)

> I cloud try JSON but with JSON we have to load all data IN-memory again before I could do any operation, with file approach we are opening the file and execution R/W operation ( we are not loading all file IN-Memory).

```
╰─$ python git_events_using_file.py
------------------------------------------------------------------------------------------------------
Top 10 active users by amount of PRs created and commits pushed (username, PRs count, commits count)

username : PR_count : commits_count
-----------------------------------
LombiqBot : 0 : 1529
renovate[bot] : 245 : 290
pull[bot] : 256 : 128
direwolf-github : 142 : 199
lihkg-boy : 0 : 331
ripamf2991 : 0 : 311
renovate-bot : 0 : 232
otiny : 0 : 222
dependabot[bot] : 149 : 34
dependabot-preview[bot] : 123 : 32
------------------------------------------------------------------------------------------------------

------------------------------------------------------------------------------------------------------
Top 10 repositories by amount of commits pushed (repo name, commits count)

repo_name : commits_count
-------------------------
lihkg-backup/thread : 331
otiny/up : 222
ripamf2991/ntdtv : 167
ripamf2991/djy : 139
wessilfie/wessilfie.github.io : 108
Lombiq/Orchard : 96
himobi/hotspot : 90
wigforss/java-8-base : 87
geos4s/18w856162 : 79
SmartThingsCommunity/SmartThingsPublic : 68
pequet/public-logs : 64
------------------------------------------------------------------------------------------------------

------------------------------------------------------------------------------------------------------
Top 10 repositories by amount of watch events (repo name, watch events count)

repo_name : watch_count
-------------------------
231135514 victorqribeiro/isocity : 44
45307548 neutraltone/awesome-stock-resources : 11
163591278 GitHubDaily/GitHubDaily : 11
184520105 sw-yx/spark-joy : 10
206874323 imsnif/bandwhich : 8
122809300 Chakazul/Lenia : 7
23560214 BurntSushi/xsv : 7
229937482 neeru1207/AI_Sudoku : 6
230327376 ErikCH/DevYouTubeList : 6
91573538 testerSunshine/12306 : 6
------------------------------------------------------------------------------------------------------
```

```
╰─$ python git_events_in-memory.py
------------------------------------------------------------------------------------------------------
Top 10 active users by amount of PRs created and commits pushed (username, PRs count, commits count)

username : PR_count : commits_count
-----------------------------------
LombiqBot : 0 : 1529
renovate[bot] : 245 : 290
pull[bot] : 256 : 128
direwolf-github : 142 : 199
lihkg-boy : 0 : 331
ripamf2991 : 0 : 311
renovate-bot : 0 : 232
otiny : 0 : 222
dependabot[bot] : 149 : 34
dependabot-preview[bot] : 123 : 32
------------------------------------------------------------------------------------------------------

------------------------------------------------------------------------------------------------------
Top 10 repositories by amount of commits pushed (repo name, commits count)

repo_name : commits_count
-------------------------
lihkg-backup/thread : 331
otiny/up : 222
ripamf2991/ntdtv : 167
ripamf2991/djy : 139
wessilfie/wessilfie.github.io : 108
Lombiq/Orchard : 96
himobi/hotspot : 90
wigforss/java-8-base : 87
geos4s/18w856162 : 79
SmartThingsCommunity/SmartThingsPublic : 68
------------------------------------------------------------------------------------------------------

------------------------------------------------------------------------------------------------------
Top 10 repositories by amount of watch events (repo name, watch events count)

repo_name : watch_count
-------------------------
231135514 victorqribeiro/isocity : 44
45307548 neutraltone/awesome-stock-resources : 11
163591278 GitHubDaily/GitHubDaily : 11
184520105 sw-yx/spark-joy : 10
206874323 imsnif/bandwhich : 8
122809300 Chakazul/Lenia : 7
23560214 BurntSushi/xsv : 7
229937482 neeru1207/AI_Sudoku : 6
230327376 ErikCH/DevYouTubeList : 6
91573538 testerSunshine/12306 : 6
------------------------------------------------------------------------------------------------------
```
