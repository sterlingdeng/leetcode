import math


def smallestValue(arr):
    left = 0
    right = len(arr)-1
    smallest = math.inf
    while left <= right:
        mid = left + int((right-left)/2)
        if arr[mid] > arr[right]:
            left = mid + 1
        else:
            smallest = min(smallest, arr[mid])
            right = mid - 1
    return smallest


print(smallestValue([6, 7, 9, 15, 19, 3, 4, 5]))
