<template>
    <a-layout id="basic-layout">
        <a-layout-sider breakpoint="sm" :trigger="null" collapsible v-model="collapsed" collapsedWidth="0">
            <div class="logo"></div>
            <a-menu :defaultSelectedKeys="[1]" mode="inline" theme="dark" @click="itemClick">
                <template  v-for="item in list">
                    <a-menu-item v-if="!item.child" :key="item.key">
                        <a-icon type="pie-chart" />
                        <span>{{item.title}}</span>
                    </a-menu-item>
                    <a-sub-menu v-else :key="item.key">
                        <span slot="title"><a-icon type="setting" /><span>{{item.title}}</span></span>
                        <a-menu-item v-for="child in item.child" :key="child.key">
                            <router-link :to="child.link">{{child.title}}</router-link>
                        </a-menu-item>
                    </a-sub-menu>
                </template>
            </a-menu>
        </a-layout-sider>
        <a-layout>
            <a-layout-header style="background: #fff; padding: 0">
                <a-icon class="trigger" :type="collapsed ? 'menu-unfold' : 'menu-fold'" @click="()=> collapsed = !collapsed"/>
            </a-layout-header>
            <a-layout-content :style="{ margin: '10px', background: '#fff', minHeight: '280px', minWidth:'350px',overflowX:'hidden'}">
                <router-view></router-view>
            </a-layout-content>
        </a-layout>
    </a-layout>
</template>
<style>
    #basic-layout {
        height: 100%;
    }

    #basic-layout .trigger {
        font-size: 18px;
        line-height: 64px;
        padding: 0 24px;
        cursor: pointer;
        /*transition: color .3s;*/
    }

    #basic-layout .trigger:hover {
        color: #1890ff;
    }

    #basic-layout .logo {
        height: 32px;
        background: rgba(255, 255, 255, .2);
        margin: 16px;
    }
</style>


<script>
    import {Layout, Menu, Icon} from 'ant-design-vue'

    export default {
        name: "BasicLayout",
        components: {
            [Layout.name]: Layout,
            [Layout.Header.name]: Layout.Header,
            [Layout.Content.name]: Layout.Content,
            [Layout.Footer.name]: Layout.Footer,

            [Layout.Sider.name]: Layout.Sider,

            [Icon.name]: Icon,
            [Menu.name]: Menu,
            [Menu.Item.name]: Menu.Item,
            [Menu.name]: Menu,
            [Menu.SubMenu.name]: Menu.SubMenu,

        },
        data() {
            return {
                collapsed: false,
                defaultNav: ['3'],
                list: [
                    {
                        key: 1,
                        icon: "dashboard",
                        title: "控制中心",
                        child: [
                            {
                                key: 2,
                                title: "工作台",
                                link:"/dashboard/workplace"
                            }
                        ]
                    },
                    {
                        key: 20,
                        icon: "dashboard",
                        title: "配置中心",
                        child: [
                            {
                                key: 22,
                                title: "分类配置",
                                link:"/dashboard/config/category"
                            },
                            {
                                key: 23,
                                title: "标签配置",
                                link:"/dashboard/config/tag"
                            },
                            {
                                key: 24,
                                title: "服务器配置",
                                link:"/dashboard/config/server"
                            }
                        ]
                    },
                    {
                        key: 40,
                        icon: "dashboard",
                        title: "媒体库",
                        child: [
                            {
                                key: 41,
                                title: "添加推广商品",
                                link:"/dashboard/add-commission-product"
                            }
                        ]
                    },
                    {
                        key: 80,
                        icon: "dashboard",
                        title: "机器人",
                        child: [
                            {
                                key: 81,
                                title: "服务器",
                                link:"/dashboard/robot/server"
                            }
                        ]
                    }
                ],
            }
        },
        methods:{
            itemClick(){
                /*console.log(this.collapsed)
                if (this.collapsed){
                    this.collapsed = false
                }*/
            },
        }
    }
</script>
