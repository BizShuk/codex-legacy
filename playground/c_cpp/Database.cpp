#include <iostream>
#include "Database.h"
#include "Composer.h"
#include <cmath>

Database::Database() {
    next_slot_ = 0;
}

Database::~Database() {
    cout << "Deconstructor of Database" << endl;
}

Composer& AddComposerUI(Database& db) {
    // input
    string in_first_name, in_last_name, in_genre, in_fact;
    int in_yob;

    cout << "First name: "; cin >> in_first_name;
    cout << "Last name: "; cin >> in_last_name;
    cout << "Genre: "; cin >> in_genre;
    cout << "Yob: "; cin >> in_yob;
    cout << "Fact: "; cin >> in_fact;


    return db.AddComposer(in_first_name, in_last_name, in_genre, in_yob, in_fact);
}


Composer& GetComposerUI(Database &db) {
    string last_name;
    cout << "Enter last name for search: "; cin >> last_name;
    return db.GetComposer(last_name);
}



Composer& PromotionUI(Composer& s) {

    int op=0;
    cout << endl;
    cout << "Promotion or Demotion" << endl;
    cout << "---------------------" << endl;
    cout << "1) Promotion with positive number >0" << endl;
    cout << "2) Deomotion with negative number <0" << endl;
    cin >> op;

    if ( op>0 ) {
        s.Promote(op);
    } else {
        s.Demote(abs(op));
    }
    s.Display();
    return s;
}

void Database::OpUI() {
    int op=0;
    Composer* s;

    cout << "Composer Database" << endl;
    cout << "----------------------------------" << endl;

    do {
        cout << "1) Add a new composer" << endl;
        cout << "2) Retrieve a composer's data" << endl;
        cout << "3) Promote/demote a composer's rank" << endl;
        cout << "4) List all composers" << endl;
        cout << "5) List all composers by rank" << endl;
        cout << "0) Quit" << endl;
        cin >> op;

        switch (op) {
            case 1: s = &AddComposerUI(*this);                      break;
            case 2: s = &GetComposerUI(*this); s->Display();        break;
            case 3: PromotionUI(*s);                                break;
            case 4: this->DisplayAll();                             break;
            case 5: this->DisplayByRank();                          break;
            default:    op = 0;                                     break;
        }
    } while ( op != 0 );

    cout << ">> Quit" << endl;
}


Composer& Database::AddComposer(string in_first_name, string in_last_name,
        string in_genre, int in_yob, string in_fact) {

    if (next_slot_ >= kMaxComposers) {
        next_slot_ = kMaxComposers - 1;
    }

    Composer* composer;
    composer = &composers_[next_slot_];

    composer->set_first_name(in_first_name);
    composer->set_last_name(in_last_name);
    composer->set_composer_yob(in_yob);
    composer->set_composer_genre(in_genre);
    composer->set_fact(in_fact);

    next_slot_++;

    return *composer;
}

Composer& Database::GetComposer(string in_last_name) {
    Composer *s;
    for ( int i = 0; i < next_slot_; ++i) {
        if (composers_[i].last_name() == in_last_name) {
            s = &composers_[i];
            return *s;
        }
    }
}

void Database::DisplayAll() {
    Composer* p;
    for (int i = 0; i < next_slot_; ++i) {
        p = &composers_[i];
        p->Display();
    }
}

void Database::DisplayByRank() {
    cout << "Sort by Rank:" << endl;
    Composer tmp;
    for (int i = 0; i < next_slot_ - 1; ++i) {
        tmp = composers_[i];
        for (int j = i+1; j < next_slot_; ++j) {
            if (tmp.ranking() < composers_[j].ranking()) {
                composers_[i] = composers_[j];
                composers_[j] = tmp;
                tmp = composers_[j];


            }
        }
    }


    this->DisplayAll();
}

