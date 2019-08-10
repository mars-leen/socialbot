<template>
    <div class="nav-list" ref="navList">
        <div class="nav-links scroll">
            <router-link v-for="(cate) in categoryList" :key="cate.Id" :to="'/media/category/'+cate.Id" class="nav-item" :style="'background-image:url('+ cate.Cover +')'">
                <span class="title">{{cate.ShortName}}</span>
            </router-link>
        </div>
    </div>
</template>

<script>
    import {listCategoryApi} from "../../api/category"
    export default {
        name: "NavList",
        data() {
            return {
                fixNavigationHeight: 82,
                categoryList : [],
            }
        },
        mounted() {
            window.addEventListener('scroll', this.handleScroll)
        },
        beforeDestroy() {
            window.removeEventListener('scroll', this.handleScroll)
        },
        created() {
            this.loadCategory()
        },
        methods: {
            handleScroll() {
                const navElement = this.$refs.navList;
                // eslint-disable-next-line no-console
                if (window.scrollY >= this.fixNavigationHeight) {
                    navElement.classList.add('fixed')
                }
                if (window.scrollY < this.fixNavigationHeight - 22) {
                    navElement.classList.remove('fixed')
                }
            },
            loadCategory() {
                listCategoryApi().then((res) => {
                    if (!res) {
                        return
                    }
                    this.categoryList = res.data
                })
            },
        }
    }
</script>

<style scoped>
    .nav-list {
        display: -webkit-box;
        display: -ms-flexbox;
        display: flex;
        -ms-flex-wrap: wrap;
        flex-wrap: wrap;
        -webkit-box-align: center;
        -ms-flex-align: center;
        align-items: center;
        -webkit-box-pack: justify;
        -ms-flex-pack: justify;
        justify-content: center;


        padding: 0.8rem 0;
        width: 100%;
        z-index: 1001;
        background: #ffffff;
    }

    .nav-list.fixed {
        position: fixed;
        width: 100%;
        box-shadow: 0 0 12px 0 hsla(0, 0%, 75.3%, .7);
        top: 0;
        padding: -0.5rem -1rem;
        max-width: 100%;
    }

    .nav-links {
        display: flex;
        align-items: center;
        justify-content: center;
        border-bottom: 1px;
    }

    .nav-list .nav-links .nav-item {
        margin-right: 1rem;
        position: relative;
        display: inline-block;
        height: 2.9rem;
        width: 5.6rem;
        font-size: 0.8rem;
        border-radius: 0.2rem;
        flex-shrink: 0;
        font-weight: 700;
        text-decoration: none;
        text-shadow: 1px 1px 0 rgba(0, 0, 0, .8);
        white-space: nowrap;
        color: #fff;

        background-position: 50%;
        background-repeat: no-repeat;
        background-size: cover;
    }

    .nav-list.fixed .nav-links .nav-item {
        height: 1.6rem;
        width: 5.5rem;
    }

    .nav-list .nav-links .nav-item .title {
        text-align: center;
        position: absolute;
        bottom: 0.2rem;
        left: 0;
        right: 0;
    }

    .nav-list.fixed .nav-links .nav-item .title {
        font-size: 0.8rem;
        bottom: 0.1rem;
    }

    @media (max-width: 576px) {
        .nav-links {
            justify-content: start;
        }

        .nav-list {
            padding: 0.6rem 0.8rem;
        }
    }


</style>
