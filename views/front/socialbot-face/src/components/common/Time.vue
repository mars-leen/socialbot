<template>
    <span>{{getTime()}}</span>
</template>

<script>
    export default {
        name: "Time",
        props:{
            time: {
                type: Number,
                default: () => {
                    return 0
                },
            },
        },
        methods:{
            getTime(){
                let minute = 1000 * 60;      //把分，时，天，周，半个月，一个月用毫秒表示
                let hour = minute * 60;
                let day = hour * 24;
                let now = new Date().getTime();   //获取当前时间毫秒
                let diffValue = now - this.time/1000000;//时间差

                if (diffValue>=0 && diffValue <= minute) {
                    return "JUST NOW"
                }
                if (diffValue > minute && diffValue <=60 *minute) {
                    return  parseInt(diffValue/minute) +" MINUTES  AGO"
                }

                if (diffValue > hour && diffValue <=24 *hour) {
                    return parseInt(diffValue/hour) +" HOURS  AGO"
                }
                if (diffValue > hour && diffValue <=24 *hour) {
                    return parseInt(diffValue/hour) +" HOURS  AGO"
                }

                return parseInt(diffValue/day) +" DAYS  AGO"
            }
        }
    }
</script>

<style scoped>

</style>