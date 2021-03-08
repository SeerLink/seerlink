import {Component, ElementRef, OnInit, ViewChild} from '@angular/core';
import {DomSanitizer} from '@angular/platform-browser';
import {HttpClient} from '@angular/common/http';
import {Base64} from 'js-base64';
import {CoinService} from '../coin/coin.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {

  constructor(public coinService:CoinService) {
  }

  ngOnInit() {
  }

  ionViewDidEnter() {

  }

  gotoScan() {
    window.open('https://scan-testnet.hecochain.com/address/'+this.coinService.lastTx+'#transactions');
  }

}
