<template>
    <a class="link-to" @click="linkTo">
        <slot></slot>
    </a>
</template>

<script>
    import {mapMutations} from "vuex"
    export default {
        name: "LinkTo",
        props:{
            to: {
                type: String,
                required: true
            },
            keep: {
                type: Boolean,
                default() {
                    return true
                },
            },
        },
        methods:{
            linkTo(){
                this.updateClickStatus(this.keep);
                this.$router.push(this.to);
            },
            ...mapMutations({
                updateClickStatus : "updateClickStatus"
            }),
        }
    }
</script>

<style scoped>
    .link-to {cursor: pointer}
</style>