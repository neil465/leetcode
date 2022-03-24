class Solution {
    public long minimumHealth(int[] damage, int armor) {
        long [][] dp = new long[damage.length][2];
        for(long [] row :dp){
            Arrays.fill(row,-1);
        }
        return minDamage(damage,0,0,armor,dp)+1;
    }
    
    private long minDamage(int [] damage,int idx,int armorUsed,int armor,long[][]dp){
        if(idx>=damage.length){
            return 0 ;
        }
        if(dp[idx][armorUsed]!=-1){
            return dp[idx][armorUsed];
        }
        if(armorUsed==0){
            long noUse = damage[idx] + minDamage(damage,idx+1,0,armor,dp);
            long use = (damage[idx]>=armor?damage[idx]-armor:0) + minDamage(damage,idx+1,1,armor,dp);
            return dp[idx][armorUsed] = Math.min(use,noUse);
        }else{
            return dp[idx][armorUsed] = damage[idx]+minDamage(damage,idx+1,1,armor,dp);
        }
    }
}