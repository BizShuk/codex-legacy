
#ifndef article_analllyzor_h
#define article_analllyzor_h 1
#include <map>
#include <list>
#include <string>

using namespace std;

const int TOKEN_SIZE = 1000;

class article_analyzor
{
public:
    article_analyzor ();
    article_analyzor ( string key , list<int> _token_list);
    ~article_analyzor ();
    void ReadArticle(string filename, int id);
    void ShowTokens();
    list<int>* gettokenlist(string key);
    article_analyzor Qand ( string key1, string key2 );
    article_analyzor Qor  ( string key1, string key2 );
private:
    map<string, list<int> > token_list;
};

#endif /* ifndef article_analllyzor_h */
