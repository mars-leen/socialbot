<template>
    <div class="layout">
        <div class="navBar">
            <nav-header></nav-header>
            <nav-list v-if="$route.meta.showNavList"></nav-list>
        </div>
        <div>
            <keep-alive>
                <router-view  v-if="$route.meta.keepAlive"></router-view>
            </keep-alive>
            <router-view v-if="!$route.meta.keepAlive"></router-view>
        </div>
    </div>
</template>

<script>
    import NavList from "../components/navbar/NavList.vue"
    import NavHeader from "../components/navbar/NavHeader.vue"
    export default {
        name: "BasicLayout",
        components: {
            NavList, NavHeader,
        }
    }
</script>

<style scoped>

</style>
