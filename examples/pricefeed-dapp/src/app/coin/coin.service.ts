import {Injectable} from '@angular/core';
import {Coin, DECIMAL} from './coin';
import {isNumber} from 'util';

var cs:CoinService;

@Injectable({
  providedIn: 'root'
})
export class CoinService {
  COINS_RUNTIME_KEY:string="COINS_RUNTIME_KEY";
  coins: Coin[];

  contractService: any;

  taskId:any;
  taskIndex:number=0;


  lastTx:string;

  constructor() {
    cs=this;

    this.initData();

    this.contractService = new ContractService();
    this.contractService.init(this.onRequestEnd);

  }

  initData() {
    this.coins = new Array();

    let sd = localStorage.getItem(this.COINS_RUNTIME_KEY);
    if (sd) {
       let info:InfoData = JSON.parse(sd) as InfoData;
       if (info) {
         this.taskIndex = info.taskIndex;
         for (let i = 0; i < info.coins.length; i++) {
           let co: Coin = new Coin();
           co.set(info.coins[i]);
           co.requestId='';
           this.coins.push(co);
         }
         return;
       }
    }

    // HT
    let coin = new Coin();
    coin.name = 'HT/USDT';
    coin.icon = 'HT.png';
    coin.base = 'HT';
    coin.quoto = 'USDT';
    coin.time = (new Date()).toLocaleString();
    coin.price = 0;
    this.coins.push(coin);

    // ETH
    coin = new Coin();
    coin.name = 'ETH/USDT';
    coin.icon = 'ETH.png';
    coin.base = 'ETH';
    coin.quoto = 'USDT';
    coin.time = (new Date()).toLocaleString();
    coin.price = 0;
    this.coins.push(coin);

    // BTC
    coin = new Coin();
    coin.name = 'BTC/USDT';
    coin.icon = 'BTC.png';
    coin.base = 'BTC';
    coin.quoto = 'USDT';
    coin.time = (new Date()).toLocaleString();
    coin.price = 0;
    this.coins.push(coin);

    // LINK
    coin = new Coin();
    coin.name = 'LINK/USDT';
    coin.icon = 'LINK.png';
    coin.base = 'LINK';
    coin.quoto = 'USDT';
    coin.time = (new Date()).toLocaleString();
    coin.price = 0;
    this.coins.push(coin);

    // ETC
    coin = new Coin();
    coin.name = 'ETC/USDT';
    coin.icon = 'ETC.png';
    coin.base = 'ETC';
    coin.quoto = 'USDT';
    coin.time = (new Date()).toLocaleString();
    coin.price = 0;
    this.coins.push(coin);

    // ADA
    coin = new Coin();
    coin.name = 'ADA/USDT';
    coin.icon = 'ADA.png';
    coin.base = 'ADA';
    coin.quoto = 'USDT';
    coin.time = (new Date()).toLocaleString();
    coin.price = 0;
    this.coins.push(coin);

    // BNB
    coin = new Coin();
    coin.name = 'BNB/USDT';
    coin.icon = 'BNB.png';
    coin.base = 'BNB';
    coin.quoto = 'USDT';
    coin.time = (new Date()).toLocaleString();
    coin.price = 0;
    this.coins.push(coin);

    // EOS
    coin = new Coin();
    coin.name = 'EOS/USDT';
    coin.icon = 'EOS.png';
    coin.base = 'EOS';
    coin.quoto = 'USDT';
    coin.time = (new Date()).toLocaleString();
    coin.price = 0;
    this.coins.push(coin);

    // NEO
    coin = new Coin();
    coin.name = 'NEO/USDT';
    coin.icon = 'NEO.png';
    coin.base = 'NEO';
    coin.quoto = 'USDT';
    coin.time = (new Date()).toLocaleString();
    coin.price = 0;
    this.coins.push(coin);

    // DOT
    coin = new Coin();
    coin.name = 'DOT/USDT';
    coin.icon = 'DOT.png';
    coin.base = 'DOT';
    coin.quoto = 'USDT';
    coin.time = (new Date()).toLocaleString();
    coin.price = 0;
    this.coins.push(coin);

    // ETH
    coin = new Coin();
    coin.name = 'ETH/JPY';
    coin.icon = 'ETH.png';
    coin.base = 'ETH';
    coin.quoto = 'JPY';
    coin.time = (new Date()).toLocaleString();
    coin.price = 0;
    this.coins.push(coin);
  }

  save() {
    let info:InfoData = new InfoData();
    info.taskIndex=this.taskIndex;
    info.coins=this.coins;
    localStorage.setItem(this.COINS_RUNTIME_KEY,JSON.stringify(info));
  }

  start() {
    if (this.taskId) {
      return;
    }
    this.onHandler();
    this.taskId=setInterval(()=>{
      this.onHandler();
    },5000);

  }

  onHandler() {
    for (let i = 0; i < this.coins.length; i++) {
      if (this.coins[i].hasRequest()) {
        return;
      }
    }

    let coin:Coin=this.coins[this.taskIndex];
    if (!coin.canRequest()) {
      return;
    }
    coin.requestId='';
    coin.requestTime=0;

    let nowD = new Date();
    console.log('Start Request:',nowD.toLocaleString(),this.taskIndex);
    this.contractService.getLastPrice(coin.base,coin.quoto,DECIMAL,this.onRequestStart,coin);
    coin.requestTime=nowD.getTime();
    this.taskIndex++;
    if (this.taskIndex >= this.coins.length) {
      this.taskIndex=0;
    }
    this.save();
  }

  stop() {
    if (this.taskId) {
      clearTimeout(this.taskId);
      this.taskId=null;
    }
  }

  onRequestStart(param:Coin,data:any) {
    if (param && data && data.events, data.events.ChainlinkRequested && data.events.ChainlinkRequested.returnValues && data.events.ChainlinkRequested.returnValues.id) {
       param.requestId=data.events.ChainlinkRequested.returnValues.id;
       console.log(data);
       cs.lastTx=data.transactionHash;
       cs.save();
    }
  }

  onRequestEnd(data:any) {
    if (!data || !data.returnValues || !data.returnValues.id) {
      return;
    }
    let requestId:string=data.returnValues.id;

    if (!requestId) {
      return;
    }
    for (let i = 0; i < cs.coins.length; i++) {
        if (cs.coins[i].requestId == requestId) {
            cs.contractService.price(cs.onGetPrice,cs.coins[i]);
        }
    }
  }

  onGetPrice(param:Coin,error:any,result:any) {
     console.log('Get Price:',param.name,param.requestId,result);
     param.requestId='';
     if (!error && result) {
       var price:number=Number(result)/DECIMAL;
       param.price=Number(price.toFixed(2));
       param.time=(new Date()).toLocaleString();
     }
     cs.save();
  }
}

class InfoData {
  taskIndex:number;
  coins: Coin[];
}
