import { Component, Inject } from '@angular/core';
import {
  MAT_LEGACY_DIALOG_DATA as MAT_DIALOG_DATA,
  MatLegacyDialogRef as MatDialogRef,
} from '@angular/material/legacy-dialog';
import { Role } from 'src/app/proto/generated/zitadel/project_pb';

@Component({
  selector: 'cnsl-user-grant-role-dialog',
  templateUrl: './user-grant-role-dialog.component.html',
  styleUrls: ['./user-grant-role-dialog.component.scss'],
})
export class UserGrantRoleDialogComponent {
  public projectId: string = '';
  public grantId: string = '';
  public selectedRoleKeysList: string[] = [];

  public selectedRoleKeys: string[] = [];

  constructor(public dialogRef: MatDialogRef<UserGrantRoleDialogComponent>, @Inject(MAT_DIALOG_DATA) public data: any) {
    this.projectId = data.projectId;
    this.grantId = data.grantId;
    this.selectedRoleKeysList = data.selectedRoleKeysList;
  }

  public selectRoles(selected: string[]): void {
    this.selectedRoleKeys = selected;
  }

  public closeDialog(): void {
    this.dialogRef.close(false);
  }

  public closeDialogWithSuccess(): void {
    this.dialogRef.close({ roles: this.selectedRoleKeys });
  }
}
