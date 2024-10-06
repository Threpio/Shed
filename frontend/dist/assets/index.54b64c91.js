(function(){const t=document.createElement("link").relList;if(t&&t.supports&&t.supports("modulepreload"))return;for(const r of document.querySelectorAll('link[rel="modulepreload"]'))l(r);new MutationObserver(r=>{for(const o of r)if(o.type==="childList")for(const u of o.addedNodes)u.tagName==="LINK"&&u.rel==="modulepreload"&&l(u)}).observe(document,{childList:!0,subtree:!0});function n(r){const o={};return r.integrity&&(o.integrity=r.integrity),r.referrerpolicy&&(o.referrerPolicy=r.referrerpolicy),r.crossorigin==="use-credentials"?o.credentials="include":r.crossorigin==="anonymous"?o.credentials="omit":o.credentials="same-origin",o}function l(r){if(r.ep)return;r.ep=!0;const o=n(r);fetch(r.href,o)}})();function E(){}function K(e){return e()}function M(){return Object.create(null)}function q(e){e.forEach(K)}function W(e){return typeof e=="function"}function G(e,t){return e!=e?t==t:e!==t||e&&typeof e=="object"||typeof e=="function"}let O;function H(e,t){return O||(O=document.createElement("a")),O.href=t,e===O.href}function J(e){return Object.keys(e).length===0}function a(e,t){e.appendChild(t)}function Q(e,t,n){e.insertBefore(t,n||null)}function z(e){e.parentNode&&e.parentNode.removeChild(e)}function m(e){return document.createElement(e)}function R(e){return document.createTextNode(e)}function x(){return R(" ")}function A(e,t,n,l){return e.addEventListener(t,n,l),()=>e.removeEventListener(t,n,l)}function i(e,t,n){n==null?e.removeAttribute(t):e.getAttribute(t)!==n&&e.setAttribute(t,n)}function U(e){return Array.from(e.childNodes)}function L(e,t){e.value=t==null?"":t}let P;function k(e){P=e}const $=[],T=[];let v=[];const F=[],V=Promise.resolve();let j=!1;function X(){j||(j=!0,V.then(D))}function I(e){v.push(e)}const S=new Set;let y=0;function D(){if(y!==0)return;const e=P;do{try{for(;y<$.length;){const t=$[y];y++,k(t),Y(t.$$)}}catch(t){throw $.length=0,y=0,t}for(k(null),$.length=0,y=0;T.length;)T.pop()();for(let t=0;t<v.length;t+=1){const n=v[t];S.has(n)||(S.add(n),n())}v.length=0}while($.length);for(;F.length;)F.pop()();j=!1,S.clear(),k(e)}function Y(e){if(e.fragment!==null){e.update(),q(e.before_update);const t=e.dirty;e.dirty=[-1],e.fragment&&e.fragment.p(e.ctx,t),e.after_update.forEach(I)}}function Z(e){const t=[],n=[];v.forEach(l=>e.indexOf(l)===-1?t.push(l):n.push(l)),n.forEach(l=>l()),v=t}const ee=new Set;function te(e,t){e&&e.i&&(ee.delete(e),e.i(t))}function ne(e,t,n,l){const{fragment:r,after_update:o}=e.$$;r&&r.m(t,n),l||I(()=>{const u=e.$$.on_mount.map(K).filter(W);e.$$.on_destroy?e.$$.on_destroy.push(...u):q(u),e.$$.on_mount=[]}),o.forEach(I)}function re(e,t){const n=e.$$;n.fragment!==null&&(Z(n.after_update),q(n.on_destroy),n.fragment&&n.fragment.d(t),n.on_destroy=n.fragment=null,n.ctx=[])}function le(e,t){e.$$.dirty[0]===-1&&($.push(e),X(),e.$$.dirty.fill(0)),e.$$.dirty[t/31|0]|=1<<t%31}function oe(e,t,n,l,r,o,u,d=[-1]){const f=P;k(e);const s=e.$$={fragment:null,ctx:[],props:o,update:E,not_equal:r,bound:M(),on_mount:[],on_destroy:[],on_disconnect:[],before_update:[],after_update:[],context:new Map(t.context||(f?f.$$.context:[])),callbacks:M(),dirty:d,skip_bound:!1,root:t.target||f.$$.root};u&&u(s.root);let c=!1;if(s.ctx=n?n(e,t.props||{},(p,h,...g)=>{const w=g.length?g[0]:h;return s.ctx&&r(s.ctx[p],s.ctx[p]=w)&&(!s.skip_bound&&s.bound[p]&&s.bound[p](w),c&&le(e,p)),h}):[],s.update(),c=!0,q(s.before_update),s.fragment=l?l(s.ctx):!1,t.target){if(t.hydrate){const p=U(t.target);s.fragment&&s.fragment.l(p),p.forEach(z)}else s.fragment&&s.fragment.c();t.intro&&te(e.$$.fragment),ne(e,t.target,t.anchor,t.customElement),D()}k(f)}class ie{$destroy(){re(this,1),this.$destroy=E}$on(t,n){if(!W(n))return E;const l=this.$$.callbacks[t]||(this.$$.callbacks[t]=[]);return l.push(n),()=>{const r=l.indexOf(n);r!==-1&&l.splice(r,1)}}$set(t){this.$$set&&!J(t)&&(this.$$.skip_bound=!0,this.$$set(t),this.$$.skip_bound=!1)}}const se="/assets/logo-lugi-white.87b9fb4a.png";function ue(e){let t,n,l,r,o,u,d,f,s,c,p,h,g,w,_,N,B;return{c(){t=m("main"),n=m("img"),r=x(),o=m("div"),o.textContent=`${ce}`,u=x(),d=m("div"),f=m("button"),f.textContent="Create",s=x(),c=m("input"),p=x(),h=m("div"),g=m("button"),g.textContent="Search",w=x(),_=m("input"),i(n,"alt","Wails logo"),i(n,"id","logo"),H(n.src,l=se)||i(n,"src",l),i(n,"class","svelte-1q31wbe"),i(o,"class","result svelte-1q31wbe"),i(o,"id","result"),i(f,"class","btn svelte-1q31wbe"),i(c,"autocomplete","off"),i(c,"class","input svelte-1q31wbe"),i(c,"id","create"),i(c,"type","text"),i(d,"class","input-box svelte-1q31wbe"),i(d,"id","inputcreate"),i(g,"class","btn svelte-1q31wbe"),i(_,"autocomplete","off"),i(_,"class","input svelte-1q31wbe"),i(_,"id","search"),i(_,"type","text"),i(h,"class","input-box svelte-1q31wbe"),i(h,"id","inputsearch")},m(b,C){Q(b,t,C),a(t,n),a(t,r),a(t,o),a(t,u),a(t,d),a(d,f),a(d,s),a(d,c),L(c,e[0]),a(t,p),a(t,h),a(h,g),a(h,w),a(h,_),L(_,e[1]),N||(B=[A(f,"click",e[2]),A(c,"input",e[4]),A(g,"click",e[3]),A(_,"input",e[5])],N=!0)},p(b,[C]){C&1&&c.value!==b[0]&&L(c,b[0]),C&2&&_.value!==b[1]&&L(_,b[1])},i:E,o:E,d(b){b&&z(t),N=!1,q(B)}}}let ce="Please enter the competition name below \u{1F447}";function ae(e,t,n){let l,r;function o(){alert(`Create Input: ${l}`)}function u(){alert(`Search Input: ${r}`)}function d(){l=this.value,n(0,l)}function f(){r=this.value,n(1,r)}return[l,r,o,u,d,f]}class fe extends ie{constructor(t){super(),oe(this,t,ae,ue,G,{})}}new fe({target:document.getElementById("app")});
