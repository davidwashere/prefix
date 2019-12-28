# Prefix
Utility to consolidate files from multiple directories into one

It will add a prefix `D` + number (ie: D0.file0.txt, D1.file0.txt) to all files, per source directory, to prevent naming collisions

Primarily used after ripping legally purchased TV shows to make automatic detection of season + episode number easier


## Install
```
git clone <this repo>
cd prefix
go install
```

## Usage
```
$ prefix
```

## Example
```
$ prefix



0. [ ] THE_SOPRANOS
1. [ ] The Sopranos Season 4 Disc 1
2. [ ] The Sopranos Season 4 Disc 2
3. [ ] The Sopranos Season 4 Disc 3
4. [ ] The Sopranos Season 4 Disc 4
5. [ ] The Sopranos Season 5 Disc 1
--------------------
Choose INPUT Dirs (seprated by spaces, nothing to continue): 1 2 3 4



0. [ ] THE_SOPRANOS
1. [X] The Sopranos Season 4 Disc 1
2. [X] The Sopranos Season 4 Disc 2
3. [X] The Sopranos Season 4 Disc 3
4. [X] The Sopranos Season 4 Disc 4
5. [ ] The Sopranos Season 5 Disc 1
--------------------
Choose INPUT Dirs (seprated by spaces, nothing to continue):



0. THE_SOPRANOS
1. The Sopranos Season 4 Disc 1
2. The Sopranos Season 4 Disc 2
3. The Sopranos Season 4 Disc 3
4. The Sopranos Season 4 Disc 4
5. The Sopranos Season 5 Disc 1
--------------------
Choose OUTPUT Dir: 0

The following renames will occur:
[The Sopranos Season 4 Disc 1/The Sopranos Season 4 Disc 1_t01.mkv] -> [THE_SOPRANOS/D0.The Sopranos Season 4 Disc 1_t01.mkv]
[The Sopranos Season 4 Disc 1/The Sopranos Season 4 Disc 1_t02.mkv] -> [THE_SOPRANOS/D0.The Sopranos Season 4 Disc 1_t02.mkv]
[The Sopranos Season 4 Disc 1/The Sopranos Season 4 Disc 1_t03.mkv] -> [THE_SOPRANOS/D0.The Sopranos Season 4 Disc 1_t03.mkv]
[The Sopranos Season 4 Disc 2/The Sopranos Season 4 Disc 2_t01.mkv] -> [THE_SOPRANOS/D1.The Sopranos Season 4 Disc 2_t01.mkv]
[The Sopranos Season 4 Disc 2/The Sopranos Season 4 Disc 2_t02.mkv] -> [THE_SOPRANOS/D1.The Sopranos Season 4 Disc 2_t02.mkv]
[The Sopranos Season 4 Disc 2/The Sopranos Season 4 Disc 2_t03.mkv] -> [THE_SOPRANOS/D1.The Sopranos Season 4 Disc 2_t03.mkv]
[The Sopranos Season 4 Disc 3/The Sopranos Season 4 Disc 3_t01.mkv] -> [THE_SOPRANOS/D2.The Sopranos Season 4 Disc 3_t01.mkv]
[The Sopranos Season 4 Disc 3/The Sopranos Season 4 Disc 3_t02.mkv] -> [THE_SOPRANOS/D2.The Sopranos Season 4 Disc 3_t02.mkv]
[The Sopranos Season 4 Disc 3/The Sopranos Season 4 Disc 3_t03.mkv] -> [THE_SOPRANOS/D2.The Sopranos Season 4 Disc 3_t03.mkv]
[The Sopranos Season 4 Disc 3/The Sopranos Season 4 Disc 3_t04.mkv] -> [THE_SOPRANOS/D2.The Sopranos Season 4 Disc 3_t04.mkv]
[The Sopranos Season 4 Disc 4/The Sopranos Season 4 Disc 4_t01.mkv] -> [THE_SOPRANOS/D3.The Sopranos Season 4 Disc 4_t01.mkv]
[The Sopranos Season 4 Disc 4/The Sopranos Season 4 Disc 4_t02.mkv] -> [THE_SOPRANOS/D3.The Sopranos Season 4 Disc 4_t02.mkv]
[The Sopranos Season 4 Disc 4/The Sopranos Season 4 Disc 4_t03.mkv] -> [THE_SOPRANOS/D3.The Sopranos Season 4 Disc 4_t03.mkv]

 13 total moves

Enter to continue, anything else cancels rename [Enter]:
Moving [The Sopranos Season 4 Disc 1/The Sopranos Season 4 Disc 1_t01.mkv] to [THE_SOPRANOS/D0.The Sopranos Season 4 Disc 1_t01.mkv]
Moving [The Sopranos Season 4 Disc 1/The Sopranos Season 4 Disc 1_t02.mkv] to [THE_SOPRANOS/D0.The Sopranos Season 4 Disc 1_t02.mkv]
Moving [The Sopranos Season 4 Disc 1/The Sopranos Season 4 Disc 1_t03.mkv] to [THE_SOPRANOS/D0.The Sopranos Season 4 Disc 1_t03.mkv]
Moving [The Sopranos Season 4 Disc 2/The Sopranos Season 4 Disc 2_t01.mkv] to [THE_SOPRANOS/D1.The Sopranos Season 4 Disc 2_t01.mkv]
Moving [The Sopranos Season 4 Disc 2/The Sopranos Season 4 Disc 2_t02.mkv] to [THE_SOPRANOS/D1.The Sopranos Season 4 Disc 2_t02.mkv]
Moving [The Sopranos Season 4 Disc 2/The Sopranos Season 4 Disc 2_t03.mkv] to [THE_SOPRANOS/D1.The Sopranos Season 4 Disc 2_t03.mkv]
Moving [The Sopranos Season 4 Disc 3/The Sopranos Season 4 Disc 3_t01.mkv] to [THE_SOPRANOS/D2.The Sopranos Season 4 Disc 3_t01.mkv]
Moving [The Sopranos Season 4 Disc 3/The Sopranos Season 4 Disc 3_t02.mkv] to [THE_SOPRANOS/D2.The Sopranos Season 4 Disc 3_t02.mkv]
Moving [The Sopranos Season 4 Disc 3/The Sopranos Season 4 Disc 3_t03.mkv] to [THE_SOPRANOS/D2.The Sopranos Season 4 Disc 3_t03.mkv]
Moving [The Sopranos Season 4 Disc 3/The Sopranos Season 4 Disc 3_t04.mkv] to [THE_SOPRANOS/D2.The Sopranos Season 4 Disc 3_t04.mkv]
Moving [The Sopranos Season 4 Disc 4/The Sopranos Season 4 Disc 4_t01.mkv] to [THE_SOPRANOS/D3.The Sopranos Season 4 Disc 4_t01.mkv]
Moving [The Sopranos Season 4 Disc 4/The Sopranos Season 4 Disc 4_t02.mkv] to [THE_SOPRANOS/D3.The Sopranos Season 4 Disc 4_t02.mkv]
Moving [The Sopranos Season 4 Disc 4/The Sopranos Season 4 Disc 4_t03.mkv] to [THE_SOPRANOS/D3.The Sopranos Season 4 Disc 4_t03.mkv]
```