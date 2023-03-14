import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { ro_RO } from 'ng-zorro-antd/i18n';
import { NzNotificationService } from 'ng-zorro-antd/notification';
import {
  FunctionsService,
  FunctionsServiceInterface,
} from 'src/app/services/functions.service';
import {
  PermalinkService,
  PermalinkServiceInterface,
} from 'src/app/services/permalink.service';

@Component({
  selector: 'app-options-permalink',
  templateUrl: './options-permalink.component.html',
  styleUrls: ['./options-permalink.component.less'],
})
export class OptionsPermalinkComponent implements OnInit {
  constructor(
    private permalinkService: PermalinkService,
    private functionsService: FunctionsService,
    private notify: NzNotificationService
  ) {}

  public havePermalink!: boolean;
  public haventPermalink!: boolean;
  public updateConfPermalink!: boolean;

  public permalink!: FunctionsServiceInterface.Receive.listPermaLink[];
  public permalinkUser!: string;

  public networks = ['testnet', 'devnet', 'mainnet'];
  public accounts!: FunctionsServiceInterface.Receive.listAccount[];

  public configPermalinkForm = new FormGroup({
    network: new FormControl(),
    recipientPermaLink: new FormControl(),
  });

  public verifyPermalink() {
    this.permalinkService.getPermalink().subscribe(
      (response) => {
        console.log(response);
        this.permalinkUser = response.url;
        this.updateConfPermalink = false;
        this.havePermalink = true;
        this.haventPermalink = false;
        this.listPermalink;
      },
      ({ error }) => {
        this.updateConfPermalink = false;
        this.havePermalink = false;
        this.haventPermalink = true;
        this.listAccount();
        this.notify.error('Error', error.error);
      }
    );
  }

  public backToPermalink() {
    this.verifyPermalink();
  }

  public updateConfigPermalink() {
    this.listAccount();
    this.havePermalink = false;
    this.haventPermalink = false;
    this.updateConfPermalink = true;
  }

  public listPermalink() {
    this.functionsService.listPermalink().subscribe(
      (response) => {
        response.forEach((link) => {
          link.expand = false;
        });

        this.permalink = response;
      },
      ({ error }) => {
        this.notify.error('Error', error.error);
      }
    );
  }

  public updatePermalink() {
    const payload = {
      userId: localStorage.getItem('id')!,
      networkPermaLink: this.configPermalinkForm.value.network,
      recipientPermaLink: this.configPermalinkForm.value.recipientPermaLink,
    };

    this.permalinkService.updatePermalink(payload).subscribe(
      (response) => {
        this.havePermalink = true;
        this.updateConfPermalink = false;
      },
      ({ error }) => {
        this.notify.error('Error', error.error);
      }
    );
  }

  public listAccount() {
    this.functionsService.listAccount().subscribe(
      (response) => {
        this.accounts = response;
      },
      ({ error }) => {
        this.notify.error('Error', error.error);
      }
    );
  }

  ngOnInit(): void {
    this.verifyPermalink();
    this.listPermalink();
  }
}
