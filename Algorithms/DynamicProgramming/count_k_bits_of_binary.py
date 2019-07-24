def countOfSubstringWithKOnes(s, k):
    N = len(s)
    res = 0
    countOfOne = 0
    freq = [0 for i in range(N + 1)]

    freq[0] = 1
 
    for i in range(0, N, 1):
        countOfOne += ord(s[i]) - ord('0')
        if (countOfOne >= k):
            res += freq[countOfOne - k]
        freq[countOfOne] += 1
    
    return res
    
def countOfSubstringWithKZeros(s, k):
    N = len(s)
    res = 0
    countOfOne = 0
    freq = [0 for i in range(N + 1)]

    freq[0] = 1
 
    for i in range(0, N, 1):
        countOfOne += ord('1') - ord(s[i])
        if (countOfOne >= k):
            res += freq[countOfOne - k]
        freq[countOfOne] += 1
    
    return res


def decimalToBinary(n): 
	return bin(n).replace("0b","") 

if __name__ == '__main__':
    n = int(input())
    for i in range(n):
        s = int(input())
        s = str(decimalToBinary(s))
        zeroval = s.count('0')
        oneval = s.count('1')
        zero = 0
        one = 0
        for i in range(1, zeroval+1, 2):
            zero += countOfSubstringWithKZeros(s, i)

        for i in range(1, oneval+1, 2):
            one += countOfSubstringWithKOnes(s, i)

        print(zero, one)