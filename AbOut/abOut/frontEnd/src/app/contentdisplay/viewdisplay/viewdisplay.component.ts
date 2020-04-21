import { Component, OnInit } from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import { HttpClient } from '@angular/common/http';
import {Outcome} from '../../models/outcome.model';
import { FormBuilder, FormGroup } from '@angular/forms';
import {Program} from '../../models/program.model';
import { ToastService } from '../../services/toast/toast.service';

@Component({
  selector: 'app-viewdisplay',
  templateUrl: './viewdisplay.component.html',
  styleUrls: ['./viewdisplay.component.scss']
})
export class ViewdisplayComponent implements OnInit {

  private outcome = {
    prefix : '',
    identifier: '',
    text: ''
  };

  private updateForm: FormGroup;
  private updateOutcomeVisible: boolean;

  private programsList: Program[] = [];
  private currentDate = new Date();
  private assocOutcomeVisible: boolean;

  constructor(private route: ActivatedRoute, 
    private http: HttpClient, 
    private toastService: ToastService,
    private formBuilder: FormBuilder,) 
    { 
      this.updateForm = this.formBuilder.group({newTextInput: ''});
    }

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      this.showOutcome(params.prefix, params.identifier);
      this.updateOutcomeVisible = false;
      this.assocOutcomeVisible = false;
    });
  }

  // getOutcome gets the JSON from the backend and stores it in the outcome object for display.
  // It also displays an error message via the toast service if one occurs.
  private showOutcome(prefix: string, identifier: string) {
    this.http.get<Outcome>('outcomes/' + prefix + '/' + identifier).subscribe(
      (response) => {
        this.outcome.identifier = response.identifier;
        this.outcome.prefix = response.prefix;
        this.outcome.text = response.text;
      },
      (error) => {
        this.toastService.show(error.status, error.message, { classname: 'bg-danger text-light' });
      }
    );
  }

  private deleteOutcome() {
    this.http.delete('outcomes/' + this.outcome.prefix + '/' + this.outcome.identifier).subscribe(
      () => {
        this.toastService.show('Successful Deletion', 'Outcome ' + this.outcome.prefix + '-' + this.outcome.identifier +
          ' successfully deleted.', { classname: 'bg-success text-light' });
      },
      (error) => {
        this.toastService.show(error.status, error.message, { classname: 'bg-danger text-light' });
      }
    )
  }

  // Sends the post request to the web service that associates the outcome.
  private associateOutcome(abbrev: string){
    this.http.post('programs/'+ abbrev + '/outcomes/' + this.outcome.prefix + "-" + this.outcome.identifier, null ).subscribe(
      () => {
        this.toastService.show('Successfully associated', 'Outcome ' + this.outcome.prefix + '-' + 
        this.outcome.identifier + ' with ' + abbrev, { classname: 'bg-success text-light' });
      },
      (error) => {
        this.toastService.show(error.status, error.message, { classname: 'bg-danger text-light' });
      }
    )
  }  
  
  private updateOutcome(description: string) {
    this.http.put('outcomes/' + this.outcome.prefix + '/' + this.outcome.identifier, description).subscribe(
      () => {
        this.toastService.show('Successfully updated', 'Outcome ' + this.outcome.prefix + '-' +
          this.outcome.identifier + ' was successfully updated', { classname: 'bg-success text-light' });
      },
      (error) => {
        this.toastService.show(error.status, error.message, { classname: 'bg-danger text-light' });
      }
    );
  }

  // getPrograms gets the JSON from the backend and stores it in programs_list.
  // It also displays an error message via the toast service if one occurs.
  private getPrograms() {
    this.http.get<Program[]>('programs').subscribe(
      (response) => {
        if (response.length !== 0) {
          this.programsList = response;
        }
      },
      (error) => {
        this.toastService.show(error.status, error.message, { classname: 'bg-danger text-light' });
      }
    );
  }

  private updateOutcomeClicked() {
    this.updateOutcomeVisible = true;
  }
}
