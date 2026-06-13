# ElasticSearch Scoring Query note

Relevence Score : Score is showing in `_score` field metadata. The higher score is, the document more related to the query.


## Boosting query
Demote query by boosting field. Take all matched positive result and multiply `revelence_scrore` with negative_boost
```es
    {
        "query": {
            "boosting": {
                "positive": { "term": { "text": "apple" } }, // must match
                "negative": { "term": { "text": "pie tart fruit crumble tree" } },
                "negative_boost": 0.5
            }
        }
    }
```

## Constant score query 
Give a constant score(boost field) when documents are filtered
```es
    {
        "query": {
            "constant_score": {
                "filter": { "term": { "user.id": "kimchy" } },
                "boost": 1.2
            }
        }
    }
```

## Disjunction max query 
Add extra score when match more. Take highest score and add with (other matched score)*`tie_breaker`
ps: dis_max will return at least one query matched resulte

```
    {
        "query": {
            "dis_max": {
                "queries": [ { "term": { "title": "Quick pets" } }, { "term": { "body": "Quick pets" } } ],
                "tie_breaker": 0.7 // floating point between 0.0 ~ 1.0. Default : 0.0.
            }
        }
    }
```


## Function Score Query
```
    {
        "query": {
            "function_score": {
                "query": { "match_all": {} },
                "boost": "5", 
                "functions": [
                    {
                        "filter": { "match": { "test": "bar" } }, 
                        "random_score": {}, 
                        "weight": 23 
                    },
                    { 
                        "filter": { "match": { "test": "cat" } }, 
                        "weight": 42
                    }
                ],
                "max_boost": 42,
                "score_mode": "max",
                "boost_mode": "multiply",
                "min_score": 42
            }
        }
    }
```

### score_mode 
function score
- `multiply` :
- `sum` :
- `avg` : Get weighted average. Example: `(score1*weight3+score2*weight4)/total_weight(3+4)`
- `first` : 
- `max` :
- `min` :

Calculated score could be limited by field `max_boost`, default `FLT_MAX`.

### boost_mode
How to combine with query score
- `multiply` query score and function score is multiplied (default)
- `replace` only function score is used, the query score is ignored
- `sum` query score and function score are added
- `avg` average
- `max` of query score and function score
- `min` of query score and function score

To exclude documents which not meet the threshold, use `min_score`


### weight
score will be multiplied with weight. 

### random_score
[0,1). No reproducible. Final score will base on `seed` field. And the minimum value of `field` for the considered document and a salt that is computed based on the index name and shard id so that documents that have the same value but are stored in different indexes get different scores. Note that documents that are within the same shard and have the same value for field will however get the same score

```
    {
        "query": {
            "function_score": {
                "random_score": {
                    "seed": 10,
                    "field": "_seq_no" // drawback for this field is that it will change when document is updated.
                }
            }
        }
    }
```

### Script score
use some value inside doc to calculate.
```
    {
        "query": {
            "function_score": {
                "query": {
                    "match": { "message": "elasticsearch" }
                },
                "script_score": {
                    "script": {
                        "params": {
                            "a": 5,
                            "b": 1.2
                        },
                        "source": "Math.log(2 + doc['my-int'].value) + params.a + params.b"
                    }
                }
            }
        }
    }
```

### Field Value factor 
Just like Script score to use value inside doc, but more simple. Should be non-negative value.

Example : `sqrt(1.2 * doc['my-int'].value)`.  But if the field is missing, use 1 instread.
```es
    {
        "query": {
            "function_score": {
                "field_value_factor": {
                    "field": "my-int",
                    "factor": 1.2,
                    "modifier": "sqrt",
                    "missing": 1
                }
            }
        }
    }
```

##### Modifier
- `none` Do not apply any multiplier to the field value
- `log` Take the common logarithm of the field value. Because this function will return a negative value and cause an error if used on values between 0 and 1, it is recommended to use log1p instead.
- `log1p` Add 1 to the field value and take the common logarithm
- `log2p` Add 2 to the field value and take the common logarithm
- `ln` Take the natural logarithm of the field value. Because this function will return a negative value and cause an error if used on values between 0 and 1, it is recommended to use ln1p instead.
- l`n1p` Add 1 to the field value and take the natural logarithm
- `ln2p` Add 2 to the field value and take the natural logarithm
- `square` Square the field value (multiply it by itself)
- `sqrt` Take the square root of the field value
- `reciprocal` Reciprocate the field value, same as 1/x where x is the field’s value



### Decay functions

check [here](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-function-score-query.html#function-decay)