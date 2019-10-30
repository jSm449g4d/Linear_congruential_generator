#include <stdio.h>
#include <time.h>
int main(){
    //X=(A*X+C)%M
    long x=0,a=214013,c=2531011,m=0x80000000-1,j=0;
    clock_t t=clock();
    x=c;
    while(x!=0){x=(a*x+c)%m;j++;}
    printf("%lf[s]:%ld[1]\n",(double)(clock()-t)/CLOCKS_PER_SEC,j);
    return 0;
}