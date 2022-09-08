## Examples

```shell
# ./csv2tsv -path /tmp/csv
# ll
-rw-r--r-- 1  501 games     195 9月   8 21:53 b.csv
-rw-r--r-- 1 root root      195 9月   8 22:47 b.tsv
-rwxr-xr-x 1  501 games 1334880 9月   8 22:45 csv2tsv
-rw-r----- 1  501 games    2991 9月   7 13:12 test_init_1.csv
-rw-r--r-- 1 root root     2991 9月   8 22:47 test_init_1.tsv
-rw-r----- 1  501 games    5007 9月   8 21:15 test_init.csv
-rw-r--r-- 1 root root     5007 9月   8 22:47 test_init.tsv
-rw-r--r-- 1  501 games 8664015 9月   8 21:52 twitter.csv
-rw-r--r-- 1 root root  8664015 9月   8 22:47 twitter.tsv
# head twitter.tsv
ItemID	Sentiment	SentimentText
1	0	                     is so sad for my APL friend.............
2	0	                   I missed the New Moon trailer...
3	1	              omg its already 7:30 :O
4	0	          .. Omgaga. Im sooo  im gunna CRy. I've been at this dentist since 11.. I was suposed 2 just get a crown put on (30mins)...
5	0	         i think mi bf is cheating on me!!!       T_T
6	0	         or i just worry too much?
7	1	       Juuuuuuuuuuuuuuuuussssst Chillin!!
8	0	       Sunny Again        Work Tomorrow  :-|       TV Tonight
9	1	      handed in my uniform today . i miss you already


# head twitter.csv
ItemID,Sentiment,SentimentText
1,0,                     is so sad for my APL friend.............
2,0,                   I missed the New Moon trailer...
3,1,              omg its already 7:30 :O
4,0,          .. Omgaga. Im sooo  im gunna CRy. I've been at this dentist since 11.. I was suposed 2 just get a crown put on (30mins)...
5,0,         i think mi bf is cheating on me!!!       T_T
6,0,         or i just worry too much?
7,1,       Juuuuuuuuuuuuuuuuussssst Chillin!!
8,0,       Sunny Again        Work Tomorrow  :-|       TV Tonight
9,1,      handed in my uniform today . i miss you already
```