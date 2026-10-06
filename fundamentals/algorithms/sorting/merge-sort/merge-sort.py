def merge_sort(array) :
    # Base case
    if not array or len(array) <= 1 : return array

    # 1. Split the array
    middle = len(array) // 2
    left = array[:middle]
    right = array[middle:]

    # 2. Recursivity
    left = merge_sort(left)
    right = merge_sort(right)

    # 3. Fuse back the array
    n1, n2 = len(left), len(right)
    n = n1 + n2
    res = [0] * n
    i, j = 0, 0
    
    while(i + j < n) :
        if i < n1 and j < n2 : 
            if left[i] < right[j] :
                res[i + j] = left[i]
                i += 1
            else :
                res[i + j] = right[j]
                j += 1
        elif i < n1 :
            res[i + j] = left[i]
            i += 1
        else :
            res[i + j] = right[j]
            j += 1

    return res