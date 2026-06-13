#include <iostream>

using namespace std;


float semester_cal(int fa, int sa, int ta, int foa, int mid_exam, int final_exam, int sec_grade ) {
    float assignment_score,mid_score,final_score,participation;

    assignment_score = ( (float) (fa+sa+ta+foa) * 0.4 / 4 );
    mid_score = (float) mid_exam * 0.15;
    final_score = (float) final_exam * 0.35;
    participation = (float) sec_grade * 0.1;

    return assignment_score + mid_score + final_score + participation;
}
void enter_score() {
    int fa,sa,ta,foa,mid_exam,final_exam,sec_grade;
    float semester_score;
    cout << "Enter the score for the first assignment : "; cin >> fa;
    cout << "Enter the score for the second assignment : "; cin >> sa;
    cout << "Enter the score for the third assignment : "; cin >> ta;
    cout << "Enter the score for the fourth assignment : "; cin >> foa;
    cout << "Enter the score for the midterm : "; cin >> mid_exam;
    cout << "Enter the score for the final : "; cin >> final_exam;
    cout << "Enter the score for the section grade : "; cin >> sec_grade;
    
    semester_score = semester_cal(fa,sa,ta,foa,mid_exam,final_exam,sec_grade);
    

    cout << "Semester score is " << semester_score << endl;
}


int main(int argc, char *argv[]){
    enter_score();
    return 0;
}
