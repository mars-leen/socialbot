<template>
    <div class="videoBoxDiv">
        <video :id="palyId" controls="controls" v-show="isVideoShow && !isloader"
               @play="videoPlay()"
               @ended="endVideo()"
               playsinline
               x-webkit-airplay="true"
               webkit-playsinline="true"
               :src="videoSrc"
        >
        </video>
        <div class="loader" v-show="isloader">
            <!--<div class="pacman"><div></div><div></div><div></div><div></div><div></div></div>-->
            <div class="loading">
                <div class="cn lt"></div>
                <div class="cn rt"></div>
                <div class="cn lb"></div>
                <div class="cn rb"></div>
            </div>
        </div>
        <img v-show="!isVideoShow " @click="palyVideo()" class="play" src="../../assets/img/palyBtn.png"/>
        <img v-show="!isVideoShow" class="loadImg" :src="loadImgUrl"/>
    </div>
</template>

<script>
    export default {
        name: 'video-play',
        props: {
            videoUrl: {
                type: String,
                defalut: ''
            },
            loadImgUrl: {
                type: String,
                defalut: ''
            },
            palyId: {
                type: String,
                defalut: ''
            }
        },
        data () {
            return {
                isVideoShow: false,
                isStop: false,
                videoSrc: '',
                isloader: false
            }
        },
        mounted () {
        },
        methods: {
            palyVideo () {
                this.isVideoShow = true
                this.isloader = true
                // 点击的时候再加载 video的
                let videoId = this.palyId
                let vdo = document.getElementById(videoId)
                let relf = this
                // 播放完和开始播放，如果有src的值就直接播放，不再重新加载一次
                if (this.videoSrc) {
                    this.isloader = false
                    vdo.play()
                } else {
                    // this.videoSrc = this.videoUrl   // IOS加载时出问题
                    vdo.setAttribute('src', this.videoUrl)
                    // vdo.load()   // ios 不识别此方法，用paly直接就有视频请求在加载了
                    // 在加载中的视频都有一个promise可以拿到 但是promise在IOS低版本上有兼容问题，直接用的监听canplay方法
                    vdo.play()
                    let canplayF = function() {
                        vdo.play()
                        relf.isloader = false
                        // vdo.removeEventListener('canplay', canplayF, false)
                    }
                    vdo.addEventListener('canplay', canplayF, false)
                    // // 检测视频加载错误时，重新加载
                    let errorF = function () {
                        vdo.load()
                        relf.isloader = false
                        relf.isVideoShow = false
                        vdo.removeEventListener('error', errorF, false)
                        vdo.removeEventListener('canplay', canplayF, false)
                    }
                    vdo.addEventListener('error', errorF, false)
                    // vdo.removeEventListener('error', errorF, false)
                }
                this.isStop = false
            },
            endVideo () {
                this.isVideoShow = false
            },
            videoPlay () {
                this.isStop = false
            }
        }
    }
</script>