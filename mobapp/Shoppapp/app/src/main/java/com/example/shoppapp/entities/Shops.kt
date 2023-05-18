package com.example.shoppapp.entities

import androidx.room.ColumnInfo
import androidx.room.Entity
import androidx.room.PrimaryKey

@Entity(tableName = "Notes")
data class Notes(

    @PrimaryKey(autoGenerate = true)
    var id:Int,

    @ColumnInfo(name = "Title")
    var title:String,

    @ColumnInfo(name = "sub_title")
    var subTitle:String,

    @ColumnInfo(name = "data_time")
    var dataTime:String,

    @ColumnInfo(name = "note_text")
    var noteText:String,

    @ColumnInfo(name = "img_path")
    var imgPath:String,

    @ColumnInfo(name = "web_link")
    var webLink:String,

    @ColumnInfo(name = "color")
    var color:String

) {
    override fun toString(): String {
        return  "$title : $dataTime"

    }
}