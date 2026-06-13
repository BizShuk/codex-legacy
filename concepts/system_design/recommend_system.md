# Recommend System
underneath information filtering system category

### important 
- recommend fresh content with preferences
- evaluation is important to make sure strategies are on the right way
- find user's need whether she expects or not.
- "The risk in recommender systems is the possibility to disturb or to upset the user which leads to a bad answer of the user"

### Strategy
- **Collaborative filtering** (also called **CF**) , use users previous experience to predict similar decision , assumption people agree in the past and will agree in the future. 
    - item-to-item CF, buy a also buy b
    - memory-based CF => User-based algorithm
    - model-based CF => Kernel-Mapping Recommender
lack:
- [Cold start](concept.md)
- scalability , large computing capability 
- sparsity , user only react on small subset of database

- **Content-based filtering** , discrete characteristics of items to get simular properties item.  but limit the content on the same type. 
    +  information retrieval
    +  information filtering

lack:
- few information on items

- **knowledge-based CF** , ?

lack:
- knowledge-based

- **Demographic**

### factor fo accuracy
- Diversity
- Recommender persistence
- Privacy
- User demographics
- Robustness
- Serendipity
- Trust
- Labelling
