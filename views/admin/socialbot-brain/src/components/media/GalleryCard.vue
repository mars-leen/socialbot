<template>
    <a-card class="card">
        <img alt="example" class="relative-img" :src="media.Url" slot="cover"/>
        <div class="card-bottom">
            <div class="cate">
                <a-radio-group v-model="activeCategory">
                    <a-radio v-for="c in categoryList" :key="c.Id" :value="c.Id">{{c.ShortName}}</a-radio>
                </a-radio-group>
            </div>
            <div class="tag">
                <span :class="tagClass(tag.Id)" v-for="tag in tags " :key="tag.Id" @click="toggleTag(tag)">
                    {{tag.ShortName}}
                </span>
            </div>
            <div class="submit">
                <a-button :loading="addLoading" class="btn" @click="addTag">发布</a-button>
            </div>
        </div>
    </a-card>

</template>

<script>
    import {Button, Tag, Select, Radio, Card} from 'ant-design-vue'
    import {addGalleryTagApi} from "../../api/gallery";

    export default {
        name: "GalleryCard",
        props: {
            categoryList: {
                type: Array,
                default: () => {
                    return []
                }
            },
            media: {
                type: Object,
                default: () => {
                    return {
                        Id: 0,
                        Cid: 0,
                        Url: "",
                        SourceType: 0,
                        Tags:[]
                    }
                }
            },
        },
        components: {
            [Card.name]: Card,
            [Radio.name]: Radio,
            [Radio.Group.name]: Radio.Group,
            [Button.name]: Button,
            [Select.name]: Select,
            [Select.Option.name]: Select.Option,
            [Tag.name]: Tag,
            [Tag.CheckableTag.name]: Tag.CheckableTag,
        },
        data() {
            return {
                addLoading: false,
                activeCategory: 0,
                activeTags: [],
                tags: [],
                category : []
            }
        },
        created(){
            if (this.media.Cid === 0) {
                this.activeCategory = this.categoryList[0].Id
            }else {
                this.activeCategory = this.media.Cid;
            }
            this.tagList()
        },
        methods: {
            toggleTag(tag) {
                const index = this.activeTags.findIndex((t) => {
                    return t === tag.Id;
                });
                if (index >= 0) {
                    this.activeTags.splice(index, 1);
                } else {
                    this.activeTags.push(tag.Id);
                }
            },
            tagClass(tagId) {
                let isSelected = this.activeTags.includes(tagId);
                return {'active': isSelected}
            },
            tagList() {
                const index = this.categoryList.findIndex((cate) => {
                    return cate.Id === this.activeCategory;
                });
                let nowCate = this.categoryList[index];
                this.tags = nowCate.Tags;

                console.log(this.media.Tags, 2222222222);
                if (this.media.Tags.length > 0) {
                    for (let i in this.media.Tags) {
                        this.activeTags.push(this.media.Tags[i].Id)
                    }
                }

            },
            addTag(){
                if (this.addLoading) {
                    return
                }
                this.addLoading = true
                const f = new FormData
                f.append("msid", this.media.Id)
                f.append("cid", this.activeCategory)
                f.append("tags",this.activeTags)
                addGalleryTagApi(f).then(res=>{
                    this.addLoading = false;
                })
            }
        }
    }
</script>

<style scoped>
    .card {border: 2px solid #ffffff;}
    .card-bottom{padding: 10px;}
    .relative-img {
        font-size: 0;
        vertical-align: top;
        line-height: 0;
        display: block;
        width: 100%;
    }
    .tag span {
        display: inline-block;
        cursor: pointer;
        margin-right: 0.3rem;
        background: #f0f0f0;
        border-radius: 0.8rem;
        padding: 0 0.5rem;
        font-size: 0.8rem;
        margin-bottom: 0.2rem;
    }

    .tag span.active {
        color: #ff6a1e;
    }
</style>
