<template>
        <div class="detail">
            <nav-bar  :is-show-nav="false"></nav-bar>
            <div class="container-min">
                <van-skeleton title :row="18" :loading="loadingDetail">
                    <h2 v-if="mediaData.Title" class="title">
                        {{mediaData.Title}}
                    </h2>go tool pprof
                    <div class="media">
                        <template v-for="(m,index) in mediaData.Medias" >
                            <div v-if="m.SourceType === 0" class="media-item" :key="index">
                                <van-image class="img"  lazy-load  fit="contain" :src="m.Url" ></van-image>
                                <van-button v-if="mediaData.ComProduct.Link !== ''" class="buy-now" type="danger" size="small" :url="mediaData.ComProduct.Link">Buy Now</van-button>
                            </div>
                            <div v-if="m.SourceType === 1" class="media-item" :key="index">
                                <video class="video"  :src="m.Url" controls="controls"></video>
                                <van-button v-if="mediaData.ComProduct.Link !== ''" class="buy-now" type="danger" size="small" :url="mediaData.ComProduct.Link">Buy Now</van-button>
                            </div>
                        </template>
                    </div>
                    <div class="extra">
                    <span class="item">
                        <van-icon name="like-o"></van-icon> <span class="small">{{mediaData.LikeNum}}</span>
                    </span>
                        <span class="item">
                        <van-icon name="fire-o"></van-icon> <span class="small">{{mediaData.ViewNum}}</span>
                     </span>
                        <span class="item">
                        <van-icon name="underway-o"></van-icon>
                        <time-print class="small" :time="mediaData.PublishAt"></time-print>
                    </span>
                    </div>
                    <div class="action">
                        <van-icon class="like active" :name="likeState()" @click="likeAction"></van-icon>
                    </div>
                </van-skeleton>
            </div>
        </div>
</template>

<script>
    import NavBar from "../../components/navbar/NavBar"
    import TimePrint from "../../components/common/Time"
    import {Tag, Icon, Skeleton, Image, Loading, Button} from 'vant';

    import {mediaDetailApi, likeMediaApi} from "../../api/media";
    import {mapGetters,mapMutations} from 'vuex'

    export default {
        name: "Detail",
        components: {
            // eslint-disable-next-line vue/no-unused-components
            NavBar, TimePrint, [Tag.name]: Tag, [Icon.name]: Icon,[Skeleton.name]:Skeleton,[Image.name]: Image, [Loading.name]: Loading, [Button.name]: Button
        },
        props:{
            uri:{
                Type:String,
                default:()=>{
                    return ""
                }
            }
        },
        data() {
            return {
                loadingDetail:true,
                isLike: false,
                mediaData: {},
                isClickTag: false,
            }
        },
        created() {
            this.onLoad();
        },
        beforeRouteLeave(to, from , next){
            to.meta.keepAlive = !this.getClickStatus();
            next();
        },
        methods: {
            ...mapGetters({
                getClickStatus: "getClickStatus"
            }),
            ...mapMutations({
                resetSelectedTags: "resetSelectedTags",
                updateClickStatus: "updateClickStatus"
            }),
            onLoad() {
                window.scrollTo(0, 0);
                mediaDetailApi({uri:this.uri}).then((res) => {
                    this.loadingDetail= false;
                    if (!res) {
                        return
                    }
                    this.mediaData = res.data;
                    this.isLike = res.data.IsLike;
                    document.title = res.data.Title;
                })
            },
            likeAction() {
                this.changeLikeShow();
                likeMediaApi(this.uri, this.isLike).then((res) => {
                    if (!res) {
                        this.changeLikeShow();
                    }
                });
            },
            changeLikeShow(){
                this.isLike = !this.isLike;
                if (this.isLike) {
                    this.mediaData.LikeNum = this.mediaData.LikeNum + 1;
                } else {
                    this.mediaData.LikeNum =  this.mediaData.LikeNum - 1;
                }
            },
            likeState() {
                if (this.isLike) {
                    return "like"
                }
                return "like-o"
            },
        }
    }
</script>

<style scoped>
    .container-min {
        margin-bottom: 3rem;
    }

    .extra .item {
        margin-right: 0.8rem;
    }

    .extra .item .small {
        margin-left: 0.4rem;
    }

    .action {
        display: -webkit-box;
        display: -ms-flexbox;
        display: flex;

        position: relative;

        -webkit-box-orient: horizontal;
        -webkit-box-direction: normal;
        -ms-flex-direction: row;
        flex-direction: row;


        -webkit-box-pack: center;
        -ms-flex-pack: center;
        justify-content: center;

        margin: 1rem;
    }

    .action .like {
        font-size: 2rem;
        border: 1px solid #ff4949;
        padding: 1rem;
        color: #ff4949;
        box-shadow: 0 2px 15px -1px #ff4949;
        border-radius: 50%;
    }
     .media .img {width: 100%;}
     .media .video{width: 100%; height: 100%;}

    .media-item {position: relative;}
    .media-item .buy-now {position: absolute; top: 0.8rem;right:0.8rem;cursor: pointer;}
</style>
