#include <stdio.h>
#include <time.h>

/**
 * Why copyij is much faster than copyij1 ???
 * The only difference is the loop's order.
 * 
*/
void copyij(int src[2048][2048],
            int dst[2048][2048])
{

  clock_t start;
  clock_t end;
  start = clock();
  int i, j;
  for (i = 0; i < 2048; i++)
  {
    for (j = 0; j < 2048; j++)
    {
      dst[i][j] = src[i][j];
    }
  }
  end = clock();
  printf("Duration = %f\n", (double)(end - start) / CLOCKS_PER_SEC);
}

void copyij2(int src[2048][2048],
             int dst[2048][2048])
{
  clock_t start;
  clock_t end;
  start = clock();
  int i, j;
  for (j = 0; j < 2048; j++)
  {
    for (i = 0; i < 2048; i++)
    {
      dst[i][j] = src[i][j];
    }
  }
  end = clock();
  printf("Duration = %f\n", (double)(end - start) / CLOCKS_PER_SEC);
}

int src[2048][2048];
int dest[2048][2048];

int main()
{
  copyij(src, dest);
  copyij2(src, dest);
  return 0;
}