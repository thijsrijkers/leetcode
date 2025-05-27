import "fmt"

func divideArray(nums []int) bool {
    startOfLoop := 0;

    for startOfLoop < len(nums) {
        secondLoop := startOfLoop + 1;
        found := false;

        for secondLoop < len(nums) {
            if nums[startOfLoop] == nums[secondLoop] {
                nums = append(nums[:secondLoop], nums[secondLoop+1:]...)
                nums = append(nums[:startOfLoop], nums[startOfLoop+1:]...)
                found = true;
                break;
            }

            secondLoop++;
        }

        if !found {
            return false;
        }
    }

    return true;
}