bool compare(string a,string b){
    return a+b>b+a;
}

class Solution {
public:
    string largestNumber(vector<int>& nums) {
        vector<string>temp;
        for(auto it:nums){
            temp.push_back(to_string(it));
        }
        sort(temp.begin(),temp.end(),compare);

        string result;
        int n=temp.size();
        for(int i=0;i<n;i++){
            result+=temp[i];
        }
        if(result[0]=='0'){
            return "0";
        }
        else{
            return result;
        }
    }
};