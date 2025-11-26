import { Controller, Delete, Get, Post, Query, Req } from '@nestjs/common';
import { UserService } from './user.service';
import {
  CurrentUser,
  Roles,
  Serialize,
  UserFullResponseDto,
  UserQueryDto,
} from '@phanhotboy/nsv-common';

@Controller('users')
export class UserController {
  constructor(private readonly userService: UserService) {}

  @Get()
  @Roles('admin')
  @Serialize(UserFullResponseDto)
  getAllUsers(@Query() query: UserQueryDto) {
    return this.userService.queryUsers(query);
  }

  @Get('me')
  @Serialize(UserFullResponseDto)
  getMe(@CurrentUser('id') userId: string) {
    return this.userService.findUserById(userId);
  }

  @Delete(':id')
  @Roles('admin')
  deleteUser(@Req() req: Request) {
    return this.userService.deleteUser(req.params.id!);
  }
}
