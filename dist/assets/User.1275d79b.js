import{i as v,j as w,r as c,o as u,c as b,a as t,w as o,v as h,k as C,t as x,e as y,E as k,m as p}from"./vendor.12d0e4e5.js";import{_ as z,l as I,r as S}from"./index.799e1907.js";const $={setup(){let i=v({list:[],total:0}),s=v({pagenum:1,pagesize:10});const d=w(),a=["\u672A\u9009\u62E9\u3001\u672A\u8D4B\u6743","\u8D85\u7EA7\u7BA1\u7406\u5458","\u524D\u7AEF","Go","Java","\u670D\u52A1\u7AEF","\u673A\u5668\u5B66\u4E60"],r=()=>{y.get(`api/admin/uninterview?pagesize=${s.pagesize}&pagenum=${s.pagenum}`).then(e=>{console.log(e),e?e.status===200&&(i.list=e.list,i.total=e.total):k.warning({message:"\u8BF7\u5148\u70B9\u51FB\u53F3\u4E0A\u89D2\u5934\u50CF\u6388\u6743\uFF0C\u83B7\u53D6\u7528\u6237\u4FE1\u606F\uFF01",type:"warning"})})};return r(),{userInfo:i,queryInfo:s,usergroup:d,evaluate:e=>{I.setCache("evalateStudent_num",e),y.get(`api/admin/gointerview/${e}`).then(g=>{console.log(g)}),S.push("../appraise")},getuserInfo:r,getusergroup:e=>a[e],handleSizeChange:e=>{console.log(e),s.pagesize=e,r()},handleCurrentChange:e=>{console.log(e),s.pagenum=e,r()}}}},j={class:"table"},B=p("\u7B49\u5F85\u9762\u8BD5\xA0 \xA0 \xA0 \xA0 \xA0 \xA0 "),q=p("\u9762\u8BD5\u4E2D \xA0 \xA0 \xA0 \xA0 \xA0 \xA0 "),E=p("\u9762\u8BD5\u7ED3\u675F\xA0"),N=p("\u53BB\u9762ta"),V={class:"block"};function D(i,s,d,a,r,m){const n=c("el-table-column"),_=c("el-button"),f=c("el-table"),e=c("el-pagination"),g=c("el-card");return u(),b("div",j,[t(g,{class:"box-card"},{default:o(()=>[t(f,{data:a.userInfo.list,border:!0,style:{"font-size":"14px"}},{default:o(()=>[t(n,{label:"\u5B66\u53F7",prop:"student_num"}),t(n,{label:"\u59D3\u540D",prop:"student_name"}),t(n,{label:"\u90AE\u7BB1",prop:"student_email"}),t(n,{label:"\u7535\u8BDD",prop:"student_telephone"}),t(n,{label:"\u9762\u8BD5\u72B6\u6001",prop:"interview_state"},{default:o(l=>[l.row.interview_state==0?(u(),h(_,{key:0,type:"primary"},{default:o(()=>[B]),_:1})):l.row.interview_state==1?(u(),h(_,{key:1,type:"success"},{default:o(()=>[q]),_:1})):(u(),h(_,{key:2,type:"warning"},{default:o(()=>[E]),_:1}))]),_:1}),t(n,{label:"\u9009\u62E9\u65B9\u5411",prop:"choice"},{default:o(l=>[C("span",null,x(a.getusergroup(l.row.choice)),1)]),_:1}),t(n,{label:"\u8BC4\u4EF7",width:"180"},{default:o(l=>[t(_,{onClick:A=>a.evaluate(l.row.student_num)},{default:o(()=>[N]),_:2},1032,["onClick"])]),_:1})]),_:1},8,["data"]),C("div",V,[t(e,{onSizeChange:a.handleSizeChange,onCurrentChange:a.handleCurrentChange,"current-page":a.queryInfo.pagenum,"page-sizes":[1,2,5,10],"page-size":a.queryInfo.pagesize,layout:"total, sizes, prev, pager, next, jumper",total:a.userInfo.total},null,8,["onSizeChange","onCurrentChange","current-page","page-size","total"])])]),_:1})])}var G=z($,[["render",D],["__scopeId","data-v-61244824"]]);const J={components:{usertable:G},setup(){}},M={class:"user"};function T(i,s,d,a,r,m){const n=c("usertable");return u(),b("div",M,[t(n)])}var K=z(J,[["render",T]]);export{K as default};