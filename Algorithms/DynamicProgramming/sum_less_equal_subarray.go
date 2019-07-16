// Java program to find subarray having 
// maximum sum less than or equal to sum 
public class Main { 

	// To find subarray with maximum sum 
	// less than or equal to sum 
	static int findMaxSubarraySum(int arr[], 
							int n, int sum) 
	{ 
	// To store current sum and 
	// max sum of subarrays 
	int curr_sum = arr[0], max_sum = 0, start = 0; 

	// To find max_sum less than sum 
	for (int i = 1; i < n; i++) { 

		// Update max_sum if it becomes 
		// greater than curr_sum 
		if (curr_sum <= sum) 
		max_sum = Math.max(max_sum, curr_sum); 

		// If curr_sum becomes greater than 
		// sum subtract starting elements of array 
		while (curr_sum + arr[i] > sum && start < i) { 
			curr_sum -= arr[start]; 
			start++; 
		} 
		
		// Add elements to curr_sum 
		curr_sum += arr[i]; 
	} 

	// Adding an extra check for last subarray 
	if (curr_sum <= sum) 
		max_sum = Math.max(max_sum, curr_sum); 

	return max_sum; 
	} 

	// Driver program to test above function 
	public static void main(String[] args) 
	{ 
		int arr[] = { 1, 2, 3, 4, 5 }; 
		int n = arr.length; 
		int sum = 11; 

		System.out.println(findMaxSubarraySum(arr, n, sum)); 
	} 
} 
