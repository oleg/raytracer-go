#include <stdio.h>

main()
{
  int fahr, celsius;
  int step, lower, upper;
  
  lower = 0;
  upper = 300;
  step = 20;
  
  fahr = lower;
  
  while(fahr <= upper) {
    celsius = 5 * (fahr - 32) / 9;
    printf("%3d\t%6d\n", fahr, celsius);
    fahr += step;
  }
}