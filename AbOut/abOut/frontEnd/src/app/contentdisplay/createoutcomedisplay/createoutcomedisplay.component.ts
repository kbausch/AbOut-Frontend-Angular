import { Component, OnInit } from '@angular/core';
import {FormBuilder} from '@angular/forms';
import { HttpClient} from '@angular/common/http';
import { ToastService } from '../../services/toast/toast.service';
import { Identifiers } from '@angular/compiler';

@Component({
  selector: 'app-createoutcomedisplay',
  templateUrl: './createoutcomedisplay.component.html',
  styleUrls: ['./createoutcomedisplay.component.scss']
})
export class CreateoutcomedisplayComponent implements OnInit {

    private outcome = {
      Prefix: '',
      Identifier: '',
      Description: ''
    }

    outcomeModel = this.outcome;

  constructor(private http: HttpClient, private toastService: ToastService) { }

  ngOnInit() {
  }

  private createOutcome(prefix: string, identifier: string, description: string) {
    this.http.post<void>('outcomes/' + prefix + '/' + identifier, description).subscribe(
      () => {
        this.toastService.show('Successfully Created', 'Outcome ' + prefix + '-' + 
        identifier + ' was successfully created.', { classname: 'bg-success text-light' });
      },
      (error) => {
        this.toastService.show(error.status, error.message, { classname: 'bg-danger text-light' });
      }
    );
  }


  

}
