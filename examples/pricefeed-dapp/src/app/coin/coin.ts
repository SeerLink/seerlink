export const REQUEST_SPACE:number = 3*60*60*1000;
export const DECIMAL:number=100;

export class Coin {
  name:string;
  icon:string;
  price:number;
  time:string;
  base:string;
  quoto:string;
  requestId:string;
  requestTime:number;

  constructor() {

  }

  getShowPrice():string {
    return this.getSymbol()+this.price.toString();
  }

  getSymbol():string {
    switch (this.quoto) {
      case 'USDT':
        return '$';
      case 'EUR':
        return '€';
      case 'JPY':
        return '円';
    }
    return '¥';
  }

  hasRequest():boolean {
    return this.requestId != null && this.requestId.length > 0;
  }

  canRequest():boolean {
    let nowD = new Date();
    if (this.requestTime+REQUEST_SPACE > nowD.getTime()) {
      return false;
    }
    return true;
  }

  set(co:Coin) {
    this.name=co.name;
    this.icon=co.icon;
    this.price=co.price;
    this.time=co.time;
    this.base=co.base;
    this.quoto=co.quoto;
    this.requestId=co.requestId;
    this.requestTime=co.requestTime;

  }
}
