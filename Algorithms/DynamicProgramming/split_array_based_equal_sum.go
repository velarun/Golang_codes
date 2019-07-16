
//O(n)
def findSplitPoint(arr, n) : 
	# traverse array element and 
	# compute sum of whole array 
	leftSum = 0
	for i in range(0, n) : 
		leftSum += arr[i] 

	# again traverse array and 
	# compute right sum and also 
	# check left_sum equal to 
	# right sum or not 
	rightSum = 0
	for i in range(n-1, -1, -1) : 
		# add current element 
		# to right_sum 
		rightSum += arr[i] 

		# exclude current element 
		# to the left_sum 
		leftSum -= arr[i] 

		if (rightSum == leftSum) : 
			return i 

	# if it is not possible 
	# to split array into 
	# two parts. 
	return -1

# Prints two parts after 
# finding split point 
# using findSplitPoint() 
def printTwoParts(arr, n) : 
	splitPoint = findSplitPoint(arr, n) 

	if (splitPoint == -1 or splitPoint == n ) : 
		print ("Not Possible") 
		return

	for i in range (0, n) : 
		if(splitPoint == i) : 
			print ("") 
		print (arr[i], end = " ")		 

# Driver Code 
arr = [1, 2, 3, 4, 5, 5] 
n = len(arr) 
printTwoParts(arr, n) 

# This code is contributed by Manish Shaw 
# (manishshaw1) 
