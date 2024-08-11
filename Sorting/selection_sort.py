def selection_sort(arr):
    n = len(arr)
    for i in range(0, n - 1):
        minimum = i
        for j in range(i + 1, n):
            if arr[j] < arr[minimum]:
                minimum = j
        arr[i], arr[minimum] = arr[minimum], arr[i]
        print(f"Value of temp arr is {arr}")
    return arr


if __name__ == "__main__":
    arr = list(
        map(int, input(f"Enter the array elements separated by space: ").split())
    )
    print(f"Value of array before sorting is {arr}")
    selection_sort(arr)
    print(f"Value of array after sorting is {arr}")
