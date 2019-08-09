<template>
    <div class="navBar">
        <nav-header></nav-header>
        <nav-list v-if="isShowNav"></nav-list>
    </div>
</template>

<script>
    import NavList from "./NavList.vue"
    import NavHeader from "./NavHeader.vue"

    export default {
        name: "Header",
        props:{
            isShowNav: {
                type: Boolean,
                default: () => {
                    return true
                }
            },
        },
        components: {
            NavList, NavHeader,
        }
    }
</script>

<style scoped>

</style>