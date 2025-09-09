using API.Models;
using API.Services;
using Microsoft.Extensions.Options;
using MongoDB.Driver;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.Configure<UrlDatabaseSettings>(
    builder.Configuration.GetSection("MongoDB"));

builder.Services.AddScoped<IUrlService, UrlService>();

builder.Services.AddSingleton<IMongoClient>(serviceProvider =>
{
    var settings = serviceProvider.GetService<IOptions<UrlDatabaseSettings>>().Value;
    return new MongoClient(settings.ConnectionString);
});

builder.Services.AddScoped(static serviceProvider =>
{
    var client = serviceProvider.GetService<IMongoClient>();
    var settings = serviceProvider.GetService<IOptions<UrlDatabaseSettings>>().Value;
    return client.GetDatabase(settings.DatabaseName);
});
builder.Services.AddControllers();


var app = builder.Build();

// Configure the HTTP request pipeline.
if (!app.Environment.IsDevelopment())
{
    // The default HSTS value is 30 days. You may want to change this for production scenarios, see https://aka.ms/aspnetcore-hsts.
    app.UseHsts();
}

app.UseRouting();

app.UseAuthorization();

app.MapControllers();

app.Run();
