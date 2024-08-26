[[easy]]

Думал о двух указателях, но так и не придумал. По итогу получилось решение `O(nlog(n) + mlog(m) + m*log(n))`. Это сложнее, чем `O(n)`, но лучше, чем если бы я использовал встроенную функцию `slices.Index`
Какой-то гений решил двумя указателями, до которых я не допер.
```cpp
class Solution {
public:
    vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
        set<int> st;
        vector<int> v;
        int i=0, j=0;
        int n1=nums1.size(), n2=nums2.size();

        sort(nums1.begin(), nums1.end());
        sort(nums2.begin(), nums2.end());

        while(i<n1 && j<n2){
            if(nums1[i] < nums2[j]) i++;
            else if(nums1[i] > nums2[j]) j++;
            else {
                st.insert(nums1[i]);
                i++;
                j++;
            }
        }

        for(auto i: st) v.push_back(i);

        return v;
    }
};
```
Тут заодно не надо выяснять какой массив длиннее. 
В чем суть, насколько я понял. Мы фиксируем указатель в 1 или 2 массиве, а другим указателем бежим по другому массиву до тех пор, пока не найдем равный элемент. 
![[Pasted image 20240826232117.png]]
