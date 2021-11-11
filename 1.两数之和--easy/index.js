var twoSum = function (nums, target) {
    const obj = {};
    for (let i = 0; i < nums.length; i++) {
        const left = target - nums[i];
        if (obj[left] !== undefined) {
            return [obj[left], i];
        } else {
            obj[nums[i]] = i;
        }
    }
};