<template>
    <a-layout id="main">
        <a-layout-sider breakpoint="sm" :trigger="null" collapsible v-model="collapsed" collapsedWidth="0">
            <div class="logo"><img src="assets/logo.png" alt=""></div>
            <a-menu theme="dark" mode="inline" :defaultSelectedKeys="defaultNav" @select="selected">
                <a-menu-item v-for="(nav, key) in navList" :key="key">
                    <a-icon :type="nav.icon" />
                    <span>{{nav.title}}</span>
                </a-menu-item>
            </a-menu>
        </a-layout-sider>
        <a-layout class="main-content">
            <a-layout-header style="background: #fff; padding: 0">
                <a-icon class="trigger" :type="collapsed ? 'menu-unfold' : 'menu-fold'" @click="()=> collapsed = !collapsed"/>
            </a-layout-header>
            <a-layout-content class="content">
                <router-view></router-view>
            </a-layout-content>
        </a-layout>
    </a-layout>
</template>

<script>
    import { Layout, Menu, Icon} from 'ant-design-vue'
    export default {
        data() {
            return {
                collapsed: false,
                defaultNav:['3'],
                navList: {
                    "1" : {
                        icon: "dashboard",
                        title:"控制中心",
                        link: "/dashboard",
                    },
                    "2": {
                        icon: "video-camera",
                        title:"媒体中心",
                        link: "/dashboard/media",
                    },
                    "3": {
                        icon: "radar-chart",
                        title:"爬虫列表",
                        link: "/dashboard/crawler",
                    }
                }
            }
        },
        created(){
            this.$router.push(this.navList[this.defaultNav[0]].link)
        },
        components:{
            [Layout.name]: Layout,
            [Layout.Header.name]: Layout.Header,
            [Layout.Sider.name]: Layout.Sider,
            [Layout.Content.name]: Layout.Content,
            [Layout.Footer.name]: Layout.Footer,
            [Menu.name]: Menu,
            [Menu.Item.name]: Menu.Item,
            [Icon.name]: Icon
        },
        methods:{
            selected(slt){
                this.$router.push(this.navList[slt.key].link)
            }
        }
    }
</script>

<style>
    #main {
        /*height: 100%;*/
    }

    #main .trigger {
        font-size: 18px;
        line-height: 64px;
        padding: 0 24px;
        cursor: pointer;
        transition: color .3s;

    }

    #main .trigger:hover {
        color: #1890ff;
    }

    #main .logo {
        height: 32px;
        background: rgba(255,255,255,.2);
        margin: 16px;
    }

    .content {
        min-width: 375px;
        overflow-x: hidden;
    }

    /*:style="{ margin: '24px 16px', padding: '24px', background: '#fff', minHeight: '280px' }"
    }*/

    @media (min-width: 576px) {
        .content {
            margin: 1rem;
        }
    }

</style>
