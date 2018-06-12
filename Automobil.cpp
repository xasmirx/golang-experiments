#include <iostream>
using namespace std;

class Automobil {
	public:
	int kilometraza;
	int godina;
	string marka;
	static int n;

	Automobil() {n++;};
	Automobil(int, int, string, string);
	~Automobil(){n--;};

	string getSasija() {
		return sasija;
	};

	private:
	string sasija;

};

int Automobil::n = 0;

Automobil::Automobil(int _kilometraza, int _godina, string _marka, string _sasija) {
	kilometraza = _kilometraza;
	godina = _godina;
	marka = _marka;
	sasija = _sasija;
	n++;
}

int main() {
	Automobil a(200200, 2002, "Golf", "WVWABCDE");
	Automobil b(12000, 2006, "Citroen", "CTRZZXXX");

	Automobil * c = new Automobil;
	cout << "Brojno stanje automobila je: " << Automobil::n << endl;
	delete c;
	cout << "Brojno stanje automobila nakon brisanja automobila c: " << Automobil::n << endl;

	cout << "----------------------------------" << endl;

	cout << "Informacije o automobilu a !!!" << endl;
	cout << "Broj sasije automobila: " << a.getSasija() << endl;


	return 0;
}
