import {Component, OnDestroy, OnInit} from '@angular/core';
import {CoinService} from '../../coin/coin.service';
import {Coin} from '../../coin/coin';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit,OnDestroy {


  constructor(public coinService:CoinService) { }

  ngOnInit() {
    this.coinService.start();
  }

  ngOnDestroy() {

  }
}


