package com;

public class DemoService {

    private static void insertSort(int[] arr) {
        for (int i = 0; i <arr.length; i++) {
            for (int j = 0; j <=i-1 ; j++) {
                if(arr[j] > arr[i]){
                    int tmp = arr[i];
                    arr[i] = arr[j];
                    arr[j] = tmp;
                }
            }
        }
        print(arr);
    }

    private static void print(int[] arr){
        for (int i = 0; i <arr.length ; i++) {
            System.out.print(arr[i] +" ");
        }
        System.out.println(" ");
    }

    public static void main(String[] args) {
        int arr[] = new int[]{1,6,6,2,12,3,5,4,8,13};
        DemoService.insertSort(arr);
     }


}
