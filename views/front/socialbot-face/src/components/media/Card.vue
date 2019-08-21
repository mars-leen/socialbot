<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
    <div class="card-con">
        <div class="card">
            <div class="card-body">
                <router-link :to="'/media/detail/' + media.Uri">
                    <van-image lazy-load class="img" fit="cover" :src="media.Cover">
                        <template v-slot:loading>
                            <van-loading type="spinner" size="20"></van-loading>
                        </template>
                    </van-image>
                </router-link >
                <span v-if="media.MediaMum > 1" class="media-num">{{media.MediaMum}}</span>
            </div>
            <div class="card-bottom">
                <router-link class="title" :to="'/media/detail/' + media.Uri" >{{media.Title}}</router-link>
                <div class="extra">
                    <div class="item">
                        <span class="user-name"></span>
                        <span class="icon"><van-icon name="fire-o"></van-icon> {{media.ViewNum}}</span>
                    </div>
                    <div class="item">
                        <span class="time"><time-print :time="media.PublishAt"></time-print></span>
                        <span class="icon"><van-icon name="like-o"></van-icon> {{media.LikeNum}}</span>
                    </div>
                </div>
            </div>
        </div>
        <img v-if="isTopRank > 0" class="top-rank " :src="getTopRank" alt="topRank">
    </div>
</template>

<script>
    import Vue from "vue";
    import {Image, Tag, Icon, Lazyload , Loading} from 'vant';
    import TimePrint from "../../components/common/Time"
    Vue.use(Lazyload);

    export default {
        name: "Card",
        components: {
            TimePrint, [Image.name]: Image, [Tag.name]: Tag, [Icon.name]: Icon, [Tag.name]: Tag, [Loading.name]: Loading
        },
        props: {
            isTopRank: {
                type: Number,
                default: () => {
                    return 0
                }
            },
            media: {
                type: Object,
                // eslint-disable-next-line vue/require-valid-default-prop
                default: () => {
                    return {
                        Uri:"",
                        Title: "",
                        Tags: [],
                        Cover: "",
                    }
                }
            },
        },
        computed: {
            getTopRank() {
                return require('../../assets/img/top' + this.isTopRank + '.png')
            }
        }
    }
</script>

<style scoped>

    .card-con {
        position: relative;
    }
    .card {
        padding: 1rem;
        position: relative;
        background: #ffffff;
        margin-bottom: 1rem;
        box-shadow: 0 0 16px 0 hsla(0, 0%, 75.3%, .7);
        transition: box-shadow .4s ease;
        border-radius: 4px;
        min-height: 3rem;overflow: hidden;
    }

    .top-rank {
        display: inline-block;
        width: 108px;
        position: absolute;
        right: -8px;
        top: -8px;
    }

    .card-bottom {
        padding: 1rem 0;
    }

    .card-body {
        position: relative;
    }

    .media-num {
        position: absolute;
        display: inline-block;
        background: #333333;
        top: 0.3rem;
        right: 0.3rem;
        opacity:0.6;
        border-radius: 50%;
        height: 1rem;
        line-height: 1rem;
        text-align: center;
        font-size: 0.7rem;
        width: 1rem;
        color: #ffffff;
    }

    .img {
        /*解决存在img的div 底部会多出空白*/
        font-size: 0;
        vertical-align: top;
        line-height: 0;
        display: block;

        min-height: 10rem;
    }

    .extra .item {
        position: relative;
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
        justify-content: space-between;
    }

    .extra .item span {
        color: #6e6a6a;
        font-size: 0.8rem;
        display: inline-block;
    }

    .extra .item .icon {
        width: 4rem;
    }

</style>
