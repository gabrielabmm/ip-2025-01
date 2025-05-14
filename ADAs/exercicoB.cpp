#include <bits/stdc++.h>
using namespace std;

int main() {
int t, cases, n, i;
vector <int> sequence;
vector <int> novsequence;

cin >> t;

for (i=0; i < t; i++){

    cin >> cases;

    sequence.clear();

for (int j=0; j < cases; j++){

    cin >> n;

sequence.push_back(n);
  }

  int k = 0;
  int j = cases - 1;
  int vez = 1;


  while (k <= j){
if (vez == 1){
novsequence.push_back(sequence[k]);
vez = 2;
k++;
} else {
novsequence.push_back(sequence[j]);
vez = 1;
j--;

}
  }

  for (int j; j < novsequence.size(); j++){
cout << novsequence[j] << " ";

  }
  cout << endl;


  novsequence.clear();

}

   return 0;
}

