<template>
    <div>
        <nav-bar></nav-bar>
        <div class="container">
            <div class="line-label">
                <span class="label">Recommend</span>
                <!--<span class="small">More</span>-->
            </div>
            <top></top>
            <div class="line-label">
                <span class="label">Newest</span>
            </div>
            <van-list  v-model="mediaList.loading" :finished="mediaList.finished" finished-text="no more" @load="onLoad">
                <masonry :cols="{default: 3, 1200: 2, 768: 2, 576: 1}" :gutter="{default: '1rem'}">
                    <card v-for="(value, index) in mediaList.items" :key="index" :media="value"></card>
                </masonry>
            </van-list>
        </div>

    </div>
</template>

<script>
    import {List} from 'vant';
    import NavBar from "../components/navbar/NavBar"
    import Card from "../components/media/Card"
    import Top from "../components/media/Top"

    import {listMediaByCategoryApi} from "../api/media";

    export default {
        name: "Home",
        components: {
            // eslint-disable-next-line vue/no-unused-components
            NavBar,Top,Card,[List.name]: List,
        },
        data(){
            return {
                mediaList: {
                    items: [],
                    lastId: 0,
                    refreshing: false,
                    loading: false,
                    error: false,
                    finished: false
                }
            }
        },
        methods:{
            onLoad(){
                listMediaByCategoryApi({lastId:this.mediaList.lastId}).then((res)=>{
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
            }
        }
    }
</script>

<style scoped>

</style>
