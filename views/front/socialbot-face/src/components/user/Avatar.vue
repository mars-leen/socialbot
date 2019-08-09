<template>
    <span class="avatar">
        <img class="relative-img" v-if="urlStr" :src="urlStr" alt="avatar">
        <van-icon v-else name="manager"></van-icon>
    </span>
</template>

<script>
    import { Icon } from 'vant';

    export default {
        name: "Avatar",
        components: {
            [Icon.name]: Icon
        },
        props:["urlStr"]
    }
</script>

<style scoped>
    .avatar {
        display: inline-block;
        font-size: 1.4rem;
        background-color: #e1e1e1;
        border-radius: 50%;
        width: 3rem;
        height: 3rem;
        line-height: 3rem;
        text-align: center;
        color: #9f9e9e;
        cursor: pointer;
        overflow: hidden;

    }

    @media  (max-width: 576px) {
        .avatar {
            width: 2rem;
            height: 2rem;
            line-height: 2rem;
        }
    }
</style>