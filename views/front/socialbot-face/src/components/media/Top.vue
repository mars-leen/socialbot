<template>
    <div class="top">
        <div class="left">
            <media-card v-if="this.topData[0]"  class="card-item" :media="this.topData[0]" :isTopRank="1"></media-card>
        </div>
        <div class="right">
            <media-card v-if="this.topData[1]" class="card-item" :media="this.topData[1]"  :isTopRank="2"></media-card>
            <media-card v-if="this.topData[2]" class="card-item" :media="this.topData[2]"  :isTopRank="3"></media-card>
        </div>
    </div>
</template>

<script>
    import {homeRecommendApi} from "../../api/media";
    import MediaCard from "./Card"
    import {Skeleton} from 'vant';

    export default {
        name: "Top",
        data() {
            return {
                topData: []
            }
        },
        components: {
            // eslint-disable-next-line vue/no-unused-components
            MediaCard,[Skeleton.name]:Skeleton,
        },
        created() {
            this.getTop();
        },
        methods: {
            getTop() {
                homeRecommendApi().then((res) => {
                    if (!res) {
                        return
                    }
                    this.topData = res.data;
                })
            }
        }
    }
</script>

<style scoped>
    .top {
        width: 100%;
    }

    .top .left {
        width: 100%;
        position: relative;
    }

    .top .right {
        width: 100%;
        position: relative;
    }

    @media (min-width: 768px) {
        .top {
            display: -webkit-box;
            display: -ms-flexbox;
            display: flex;

            -webkit-box-orient: vertical;
            -webkit-box-direction: normal;

            -ms-flex-direction: column;
            flex-direction: column;

            -webkit-box-align: start;
            -ms-flex-align: start;
            align-items: flex-start;

            -webkit-box-pack: center;
            -ms-flex-pack: center;
            justify-content: center;
        }

        .top .left {
            height: 26rem;
        }
        .top .right {
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
            justify-content: space-between;
        }

        .top .left .card-item{
            height: 96%;
        }

        .top .right .card-item {
            margin-bottom: 0;
            -webkit-box-flex: 0;
            -webkit-flex: 0 0 49%;
            flex: 0 0 49%;
        }


    }


    @media (min-width: 1200px) {
        .top {
            height: 26rem;

            display: -webkit-box;
            display: -ms-flexbox;
            display: flex;

            -webkit-box-orient: horizontal;
            -webkit-box-direction: normal;

            -ms-flex-direction: row;
            flex-direction: row;

            -webkit-box-align: stretch;
            -ms-flex-align: stretch;
            align-items: stretch;

            -webkit-box-pack: center;
            -ms-flex-pack: center;
            justify-content: center;
        }

        .top .right {
            position: relative;

            display: -webkit-box;
            display: -ms-flexbox;
            display: flex;

            -webkit-box-orient: vertical;
            -webkit-box-direction: normal;
            -ms-flex-direction: column;
            flex-direction: column;

            -webkit-box-pack: center;
            -ms-flex-pack: center;
            justify-content: space-between;

            -webkit-box-align: stretch;
            -ms-flex-align: stretch;
            align-items: stretch;

            -webkit-box-flex: 0;
            -ms-flex: 0 0 32%;
            flex: 0 0 32%;
            max-width: 32%;
        }

        .top .left {
            -webkit-box-flex: 0;
            -ms-flex: 0 0 68%;
            flex: 0 0 68%;
            max-width: 68%;
            padding-right: 1rem;
        }

        .top .left .card-item{
            height: 100%;
        }

        .top .right .card-item {
            margin-bottom: 0;
            -webkit-box-flex: 0;
            -ms-flex: 0 0 40%;
            flex: 0 0 48%;
        }

        .top .card-item .card-bottom {
            position: relative;
            color: #ffffff !important;
            bottom:1rem;
            left:1rem;
            right: 1rem;
        }

    }


</style>
