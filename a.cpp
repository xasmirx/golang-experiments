#include <iostream>
using namespace std;

class Osoba {
	private:
	float visina;
	float tezina;
	int godine;

	public:
	void setujPodatke(float, float, int);
	float getVisina();
	float getTezina();
	int getGodine();
};

float Osoba::getVisina() {

	if(visina == 5) {
		visina = 150;
	}

	return visina;
}

float Osoba::getTezina() {
	return tezina;
}

int Osoba::getGodine() {
	return godine;
}

void Osoba::setujPodatke(float a, float b, int c) {
	visina = a;
	tezina = b;
	godine = c;
}

class Student : public Osoba {
	public:
	int godina;
	int brojIndeksa;

	Student(float, float, int, int, int);	

};

Student::Student(float a, float b, int c, int d, int e) {
	setujPodatke(a, b, c);
	cout << getVisina() << getTezina() << getGodine();
	godina = d;
	brojIndeksa = e;
}

int main() {
	Student asmir(5, 2, 3, 4, 5);
	return 1;
}


