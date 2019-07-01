#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#define SIZE 6
int v[SIZE][SIZE];
void *function(void *arg);
void *function2(void *arg);
void *criar(void *arg);
pthread_t t1,t2,t3;
int main(){
    //pthread_t t3;
    int a1=1;
    int a2=2;
    int a3=3;
    int i,j;
    while(1){
        pthread_create(&t3,NULL,criar,(void *)(&a3));
        pthread_join(t3,NULL);

        printf("\n");
        for(i=0;i < SIZE ; i++){
            for(j=0;j<SIZE;j++){
                v[i][j] = i + j;
                printf("%d",v[i][j]);
            }
            printf("\n");
        }
        sleep(1);
        pthread_exit(NULL);
    }
    exit(0);
}
void *  criar (void *arg){
    //pthread_t t1,t2;
    int i,j;
    int a1=1;
    int a2=2;
    //printf(“valor de retorno do argumento:%d \n”,*valor);
    pthread_create(&t1,NULL,function,(void *)(&a1));
    pthread_join(t1,NULL);
    pthread_exit(NULL);
};
void * function (void *arg){
    int *valor = (int *)(arg);
    int i,j;
    for(i=0; i<SIZE/2;i++){
        for(j=0; j<SIZE;j++){
            v[i][j]= i + j;
            printf("%d",v[i][j]);
        }
        printf("\n");
    }
    pthread_exit(NULL);
};