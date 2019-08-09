<template>
    <div>
        <nav-bar></nav-bar>
        <div class="container">
            <sort @click="getSort"></sort>
            <van-list class="roll-list"  v-model="mediaList.loading" :finished="mediaList.finished" finished-text="no more"
                      @load="onLoad(false)">
                <masonry :cols="{default: 3, 1204: 3, 768: 2, 576: 1}" :gutter="{default: '1rem'}">
                    <card v-for="(value, index) in mediaList.items" :key="index" :media="value"></card>
                </masonry>
            </van-list>
        </div>
    </div>
</template>

<script>
    import {List} from 'vant';
    import NavBar from "../../components/navbar/NavBar"
    import Card from "../../components/media/Card"
    import Sort from "../../components/media/Sort"

    import {listMediaByCategoryApi} from "../../api/media";

    export default {
        name: "Category",
        components: {
            NavBar, Card, Sort, [List.name]: List,
        },
        props:{
            id:{
                Type:String,
                default:()=>{
                    return "0"
                }
            }
        },
        created(){
            this.cateId = this.id
        },
        data() {
            return {
                sort:1,
                cateId:0,
                mediaList: {
                    items: [],
                    lastId: 0,
                    loading: false,
                    finished: false
                }
            }
        },
        beforeRouteUpdate (to, from, next) {
            this.cateId = to.params.id;
            this.mediaList.lastId = 0;
            this.mediaList.loading = false;
            this.mediaList.finished = false;
            this.mediaList.items=[];
            window.scrollTo(0, 10);
            next();
        },
        methods: {
            onLoad() {
                listMediaByCategoryApi({lastId:this.mediaList.lastId, category:this.cateId, sort:this.sort}).then((res) => {
                    // eslint-disable-next-line no-console
                    this.mediaList.loading = false;
                    if (!res) {
                        this.mediaList.finished = true;
                        return
                    }
                    let l = res.data.length;
                    if (l < 15) {
                        this.mediaList.finished = true;
                    }
                    if (l !== 0) {
                        this.mediaList.lastId = res.data[l - 1].LastId
                    }
                    for (let i = 0; i < l; i++) {
                        this.mediaList.items.push(res.data[i]);
                    }
                })
            },
            getSort(current){
                this.sort = current;
                this.mediaList.lastId = 0;
                this.mediaList.loading = false;
                this.mediaList.finished = false;
                this.mediaList.items=[];
                window.scrollTo(0, 10);
            },
        }
    }
</script>

<style scoped>
</style>
