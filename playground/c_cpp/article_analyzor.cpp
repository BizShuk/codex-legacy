#include <iostream>
#include "article_analyzor.h"
#include <string>
#include <map>
#include <fstream>
#include "string.h"
#include <vector>
#include <boost/algorithm/string.hpp>
#include <boost/algorithm/string/split.hpp>


using namespace std;


article_analyzor::article_analyzor()
{

}
article_analyzor::article_analyzor ( string key , list<int> _token_list )
{
    this->token_list[key] = _token_list;
}

article_analyzor::~article_analyzor()
{

}

// open file, read token distinct, and assign to token_list
void article_analyzor::ReadArticle(string filename , int id)
{
    map<string, int> tokenmap;
    string line;
    vector<string> parts;


    ifstream areader(filename);

    if (areader.fail())
    {
        cout << "File " << filename << " is not existed." << endl;
        return;
    }

    while ( areader >> line )
    {
        boost::split( parts, line, boost::is_any_of(" .,\"';:|()@") );
        for (int i = 0; i < parts.size(); ++i)
        {
            tokenmap[parts[i]] = 1;
        }
    }

    for (map<string, int>::iterator it = tokenmap.begin(); it != tokenmap.end(); ++it)
    {
        this->token_list[it->first].push_back(id);
    }

    areader.close();

}

void article_analyzor::ShowTokens()
{
    map<string, list<int> >* t = &this->token_list;
    list<int>* l;

    for ( auto it = t->begin(); it != t->end(); ++it )
    {
        cout << it->first << endl;
        l = &it->second;
        for (auto it_list = l->begin(); it_list != l->end(); ++it_list)
        {
            cout << *it_list << " " << endl;
        }

    }
}

list<int>* article_analyzor::gettokenlist(string key)
{
    map<string, list<int> >* t = &this->token_list;

    if ( t->find(key) == t->end())
    {
        return NULL;
    }
    else
    {
        return &t->find(key)->second;
    }

}

article_analyzor article_analyzor::Qand ( string key1 , string key2 )
{

    list<int>* a = this->gettokenlist(key1);
    list<int>* b = this->gettokenlist(key2);
    list<int> ab;

    auto bt = b->begin();

    for ( auto at = a->begin(); at != a->end() ; ++at)
    {
        while ( bt != b->end() && *bt < *at)
        {
            bt++;
        }

        if ( *at == *bt )
        {
            ab.push_back(*at);
        }
    }


    article_analyzor ret( key1 + " && " + key2 , ab);

    return ret;
}


article_analyzor article_analyzor::Qor ( string key1 , string key2 )
{

    list<int>* a = this->gettokenlist(key1);
    list<int>* b = this->gettokenlist(key2);
    list<int> ab;

    auto bt = b->begin();

    for ( auto at = a->begin(); at != a->end() ; ++at )
    {
        ab.push_back(*at);

        while ( bt != b->end() && *bt < *at )
        {
            ab.push_back(*bt);
            bt++;
        }

        if ( *bt == *at )
        {
            bt++;
        }
    }

    while ( bt != b->end() )
    {
        ab.insert(ab.end(), bt, b->end());
    }

    article_analyzor ret( key1 + " || " + key2 , ab);

    return ret;
}



int main(int argc, char *argv[])
{
    article_analyzor a;
    a.ReadArticle("article_1.txt", 1);
    a.ReadArticle("article_2.txt", 2);
    a.ReadArticle("article_3.txt", 3);
    a.ReadArticle("article_4.txt", 4);
    //a.ShowTokens();
    article_analyzor b = a.Qor("willing", "with");
    b.ShowTokens();



    return 0;
}
