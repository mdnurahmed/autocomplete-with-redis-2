# autocomplete-with-redis-2

A simple scalabale implimentation of search autocomplete using Redis based on the [blog of Salvatore Sanfilippo (Antirez)](http://oldblog.antirez.com/post/autocomplete-with-redis.html) , the creator of Redis . On the backend I used Golang. 

This solution can suggest top N word based on search frequency .  

[Here is a blog post explaining this method in details](https://dev.to/mdnurahmed/simple-scalable-search-autocomplete-systems-1j18)

# How To Run 
Using Docker - 
```
git clone https://github.com/mdnurahmed/autocomplete-with-redis-2
cd autocomplete-with-redis-2
docker-compose up --build
```

Then go to localhost:3000 in the browser . I have included redisinsight with the docker-compose file. So if you wanna see how the search strings are stored in redis go to localhost:8001 and connect to the Redis inside the docker network with the following credentials -

```
Host : redis 
Port : 6379 
Name: redis 
Username :
Password :
```
